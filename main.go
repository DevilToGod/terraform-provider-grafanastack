package main

import (
	"context"
	"flag"
	"log"

	"github.com/DevilToGod/terraform-provider-grafanastack/grafanastack"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)

func main() {

	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: grafanastack.Provider(version)}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/grafana/grafana", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
