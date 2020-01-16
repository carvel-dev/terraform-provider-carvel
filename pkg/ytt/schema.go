package ytt

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	schemaFilesKey                 = "files"
	schemaIgnoreUnknownCommentsKey = "ignore_unknown_comments"
	schemaValuesYAMLKey            = "values_yaml"
	schemaResultKey                = "result"
	schemaDebugLogsKey             = "debug_logs"
)

var (
	resourceSchema = map[string]*schema.Schema{
		schemaFilesKey: {
			Type:        schema.TypeList,
			Description: "Files",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		schemaIgnoreUnknownCommentsKey: {
			Type:        schema.TypeBool,
			Description: "Set to ignore unknown comments",
			Optional:    true,
			Default:     false,
		},
		schemaValuesYAMLKey: {
			Type:        schema.TypeString,
			Description: "Data values as YAML",
			Optional:    true,
			Sensitive:   true,
		},
		schemaResultKey: {
			Type:        schema.TypeString,
			Description: "Result",
			Computed:    true,
			Sensitive:   true,
		},
		schemaDebugLogsKey: {
			Type:        schema.TypeBool,
			Description: "Enable debug logging",
			Optional:    true,
			Default:     false,
		},
	}
)
