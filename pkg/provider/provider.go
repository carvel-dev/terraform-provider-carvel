package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/k14s/terraform-provider-k14s/pkg/kapp"
	"github.com/k14s/terraform-provider-k14s/pkg/logger"
	"github.com/k14s/terraform-provider-k14s/pkg/ytt"
)

func Provider() terraform.ResourceProvider {
	logger := logger.MustNewFileRoot("/tmp/terraform-provider-k14s.log")
	yttLogger := logger.WithLabel("ytt")

	// TODO different naming
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"k14s_ytt": ytt.NewResource(yttLogger),
		},
		ResourcesMap: map[string]*schema.Resource{
			"k14s_ytt":  schema.DataSourceResourceShim("k14s_ytt", ytt.NewResource(yttLogger)),
			"k14s_kapp": kapp.NewResource(logger.WithLabel("kapp")),
		},
	}
}
