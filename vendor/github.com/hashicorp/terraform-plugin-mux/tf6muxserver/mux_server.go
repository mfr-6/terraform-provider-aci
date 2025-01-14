// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf6muxserver

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/internal/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ tfprotov6.ProviderServer = &muxServer{}

// muxServer is a gRPC server implementation that stands in front of other
// gRPC servers, routing requests to them as if they were a single server. It
// should always be instantiated by calling NewMuxServer().
type muxServer struct {
	// Routing for data source types
	dataSources map[string]tfprotov6.ProviderServer

	// Routing for resource types
	resources map[string]tfprotov6.ProviderServer

	// Resource capabilities are cached during GetMetadata/GetProviderSchema
	resourceCapabilities map[string]*tfprotov6.ServerCapabilities

	// serverDiscoveryComplete is whether the mux server's underlying server
	// discovery of resource types has been completed against all servers.
	// If false during a resource type specific RPC, the mux server needs to
	// pre-emptively call the GetMetadata RPC or GetProviderSchema RPC (as a
	// fallback) so it knows which underlying server should receive the RPC.
	serverDiscoveryComplete bool

	// serverDiscoveryDiagnostics caches diagnostics found during server
	// discovery so they can be returned for later requests if necessary.
	serverDiscoveryDiagnostics []*tfprotov6.Diagnostic

	// serverDiscoveryMutex is a mutex to protect concurrent server discovery
	// access from race conditions.
	serverDiscoveryMutex sync.RWMutex

	// Underlying servers for requests that should be handled by all servers
	servers []tfprotov6.ProviderServer
}

// ProviderServer is a function compatible with tf6server.Serve.
func (s *muxServer) ProviderServer() tfprotov6.ProviderServer {
	return s
}

func (s *muxServer) getDataSourceServer(ctx context.Context, typeName string) (tfprotov6.ProviderServer, []*tfprotov6.Diagnostic, error) {
	s.serverDiscoveryMutex.RLock()
	server, ok := s.dataSources[typeName]
	discoveryComplete := s.serverDiscoveryComplete
	s.serverDiscoveryMutex.RUnlock()

	if discoveryComplete {
		if ok {
			return server, s.serverDiscoveryDiagnostics, nil
		}

		return nil, []*tfprotov6.Diagnostic{
			dataSourceMissingError(typeName),
		}, nil
	}

	err := s.serverDiscovery(ctx)

	if err != nil || diagnosticsHasError(s.serverDiscoveryDiagnostics) {
		return nil, s.serverDiscoveryDiagnostics, err
	}

	s.serverDiscoveryMutex.RLock()
	server, ok = s.dataSources[typeName]
	s.serverDiscoveryMutex.RUnlock()

	if !ok {
		return nil, []*tfprotov6.Diagnostic{
			dataSourceMissingError(typeName),
		}, nil
	}

	return server, s.serverDiscoveryDiagnostics, nil
}

func (s *muxServer) getResourceServer(ctx context.Context, typeName string) (tfprotov6.ProviderServer, []*tfprotov6.Diagnostic, error) {
	s.serverDiscoveryMutex.RLock()
	server, ok := s.resources[typeName]
	discoveryComplete := s.serverDiscoveryComplete
	s.serverDiscoveryMutex.RUnlock()

	if discoveryComplete {
		if ok {
			return server, s.serverDiscoveryDiagnostics, nil
		}

		return nil, []*tfprotov6.Diagnostic{
			resourceMissingError(typeName),
		}, nil
	}

	err := s.serverDiscovery(ctx)

	if err != nil || diagnosticsHasError(s.serverDiscoveryDiagnostics) {
		return nil, s.serverDiscoveryDiagnostics, err
	}

	s.serverDiscoveryMutex.RLock()
	server, ok = s.resources[typeName]
	s.serverDiscoveryMutex.RUnlock()

	if !ok {
		return nil, []*tfprotov6.Diagnostic{
			resourceMissingError(typeName),
		}, nil
	}

	return server, s.serverDiscoveryDiagnostics, nil
}

// serverDiscovery will populate the mux server "routing" for resource types by
// calling all underlying server GetMetadata RPC and falling back to
// GetProviderSchema RPC. It is intended to only be called through
// getDataSourceServer and getResourceServer.
//
// The error return represents gRPC errors, which except for the GetMetadata
// call returning the gRPC unimplemented error, is always returned.
func (s *muxServer) serverDiscovery(ctx context.Context) error {
	s.serverDiscoveryMutex.Lock()
	defer s.serverDiscoveryMutex.Unlock()

	// Return early if subsequent concurrent operations reached this logic.
	if s.serverDiscoveryComplete {
		return nil
	}

	logging.MuxTrace(ctx, "starting underlying server discovery via GetMetadata or GetProviderSchema")

	for _, server := range s.servers {
		ctx := logging.Tfprotov6ProviderServerContext(ctx, server)
		ctx = logging.RpcContext(ctx, "GetMetadata")

		logging.MuxTrace(ctx, "calling GetMetadata for discovery")
		metadataResp, err := server.GetMetadata(ctx, &tfprotov6.GetMetadataRequest{})

		// GetMetadata call was successful, populate caches and move on to next
		// underlying server.
		if err == nil && metadataResp != nil {
			// Collect all underlying server diagnostics, but skip early return.
			s.serverDiscoveryDiagnostics = append(s.serverDiscoveryDiagnostics, metadataResp.Diagnostics...)

			for _, serverDataSource := range metadataResp.DataSources {
				if _, ok := s.dataSources[serverDataSource.TypeName]; ok {
					s.serverDiscoveryDiagnostics = append(s.serverDiscoveryDiagnostics, dataSourceDuplicateError(serverDataSource.TypeName))

					continue
				}

				s.dataSources[serverDataSource.TypeName] = server
			}

			for _, serverResource := range metadataResp.Resources {
				if _, ok := s.resources[serverResource.TypeName]; ok {
					s.serverDiscoveryDiagnostics = append(s.serverDiscoveryDiagnostics, resourceDuplicateError(serverResource.TypeName))

					continue
				}

				s.resources[serverResource.TypeName] = server
				s.resourceCapabilities[serverResource.TypeName] = metadataResp.ServerCapabilities
			}

			continue
		}

		// Only continue if the gRPC error was an unimplemented code, otherwise
		// return any other gRPC error immediately.
		grpcStatus, ok := status.FromError(err)

		if !ok || grpcStatus.Code() != codes.Unimplemented {
			return err
		}

		logging.MuxTrace(ctx, "calling GetProviderSchema for discovery")
		providerSchemaResp, err := server.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})

		if err != nil {
			return err
		}

		// Collect all underlying server diagnostics, but skip early return.
		s.serverDiscoveryDiagnostics = append(s.serverDiscoveryDiagnostics, providerSchemaResp.Diagnostics...)

		for typeName := range providerSchemaResp.DataSourceSchemas {
			if _, ok := s.dataSources[typeName]; ok {
				s.serverDiscoveryDiagnostics = append(s.serverDiscoveryDiagnostics, dataSourceDuplicateError(typeName))

				continue
			}

			s.dataSources[typeName] = server
		}

		for typeName := range providerSchemaResp.ResourceSchemas {
			if _, ok := s.resources[typeName]; ok {
				s.serverDiscoveryDiagnostics = append(s.serverDiscoveryDiagnostics, resourceDuplicateError(typeName))

				continue
			}

			s.resources[typeName] = server
			s.resourceCapabilities[typeName] = providerSchemaResp.ServerCapabilities
		}
	}

	s.serverDiscoveryComplete = true

	return nil
}

// NewMuxServer returns a muxed server that will route gRPC requests between
// tfprotov6.ProviderServers specified. When the GetProviderSchema RPC of each
// is called, there is verification that the overall muxed server is compatible
// by ensuring:
//
//   - All provider schemas exactly match
//   - All provider meta schemas exactly match
//   - Only one provider implements each managed resource
//   - Only one provider implements each data source
func NewMuxServer(_ context.Context, servers ...func() tfprotov6.ProviderServer) (*muxServer, error) {
	result := muxServer{
		dataSources:          make(map[string]tfprotov6.ProviderServer),
		resources:            make(map[string]tfprotov6.ProviderServer),
		resourceCapabilities: make(map[string]*tfprotov6.ServerCapabilities),
		servers:              make([]tfprotov6.ProviderServer, 0, len(servers)),
	}

	for _, server := range servers {
		result.servers = append(result.servers, server())
	}

	return &result, nil
}
