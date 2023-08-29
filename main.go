package main

import (
	"context"
	"flag"
	"log"

	"github.com/CiscoDevNet/terraform-provider-aci/v2/aci"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return aci.Provider()
		},
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "github.com/ciscodevnet/aci", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
