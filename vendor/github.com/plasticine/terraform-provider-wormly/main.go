package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/plasticine/terraform-provider-wormly/wormly"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: wormly.Provider})
}
