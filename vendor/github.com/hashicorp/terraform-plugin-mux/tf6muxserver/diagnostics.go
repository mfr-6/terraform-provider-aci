// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf6muxserver

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func dataSourceDuplicateError(typeName string) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Invalid Provider Server Combination",
		Detail: "The combined provider has multiple implementations of the same data source type across underlying providers. " +
			"Data source types must be implemented by only one underlying provider. " +
			"This is always an issue in the provider implementation and should be reported to the provider developers.\n\n" +
			"Duplicate data source type: " + typeName,
	}
}

func dataSourceMissingError(typeName string) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Data Source Not Implemented",
		Detail: "The combined provider does not implement the requested data source type. " +
			"This is always an issue in the provider implementation and should be reported to the provider developers.\n\n" +
			"Missing data source type: " + typeName,
	}
}

func diagnosticsHasError(diagnostics []*tfprotov6.Diagnostic) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic == nil {
			continue
		}

		if diagnostic.Severity == tfprotov6.DiagnosticSeverityError {
			return true
		}
	}

	return false
}

func resourceDuplicateError(typeName string) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Invalid Provider Server Combination",
		Detail: "The combined provider has multiple implementations of the same resource type across underlying providers. " +
			"Resource types must be implemented by only one underlying provider. " +
			"This is always an issue in the provider implementation and should be reported to the provider developers.\n\n" +
			"Duplicate resource type: " + typeName,
	}
}

func resourceMissingError(typeName string) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Resource Not Implemented",
		Detail: "The combined provider does not implement the requested resource type. " +
			"This is always an issue in the provider implementation and should be reported to the provider developers.\n\n" +
			"Missing resource type: " + typeName,
	}
}
