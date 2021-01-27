package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.Provider})
}
