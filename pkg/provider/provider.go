package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/kapp"
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/kbld"
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/logger"
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/schemamisc"
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/ytt"
)

func Provider() terraform.ResourceProvider {
	logger := logger.MustNewFileRoot("/tmp/terraform-provider-carvel.log")
	yttLogger := logger.WithLabel("ytt")
	kbldLogger := logger.WithLabel("kbld")

	// TODO different naming
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"carvel_ytt":  ytt.NewResource(yttLogger),
			"carvel_kbld": kbld.NewResource(kbldLogger),
		},
		ResourcesMap: map[string]*schema.Resource{
			"carvel_ytt":  schema.DataSourceResourceShim("carvel_ytt", ytt.NewResource(yttLogger)),
			"carvel_kbld": schema.DataSourceResourceShim("carvel_kbld", kbld.NewResource(kbldLogger)),
			"carvel_kapp": kapp.NewResource(logger.WithLabel("kapp")),
		},

		Schema: resourceSchema,

		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return schemamisc.Context{
				Kubeconfig:  NewKubeconfig(d),
				DiffPreview: kappDiffPreviewValue(d),
			}, nil
		},
	}
}
