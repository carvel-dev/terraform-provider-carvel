package ytt

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	schemaFilesKey                 = "files"
	schemaIgnoreUnknownCommentsKey = "ignore_unknown_comments"
	schemaValuesYAMLKey            = "values_yaml"
	schemaValuesKey                = "values"
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
		schemaValuesKey: {
			Type:        schema.TypeMap,
			Description: "Data values as strings",
			Optional:    true,
			Default:     map[string]interface{}{},
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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
