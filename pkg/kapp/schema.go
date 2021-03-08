package kapp

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	schemaAppKey       = "app"
	schemaNamespaceKey = "namespace"

	schemaConfigYAMLKey = "config_yaml"
	schemaFilesKey      = "files"

	schemaDiffChangesKey = "diff_changes"
	schemaDiffContextKey = "diff_context"

	schemaDebugLogsKey = "debug_logs"

	schemaDeployKey     = "deploy"
	schemaDeleteKey     = "delete"
	schemaRawOptionsKey = "raw_options"

	// Fields automatically managed by the kapp resource
	// (Needed two different diff previews because computed property
	// updates were not shown during update operation)
	schemaClusterDriftDetectedKey = "cluster_drift_detected"
	schemaDiffPreview1Key         = "diff_preview_1"
	schemaDiffPreview2Key         = "diff_preview_2"
)

var (
	resourceSchema = map[string]*schema.Schema{
		schemaAppKey: {
			Type:        schema.TypeString,
			Description: "App name",
			Required:    true,
			ForceNew:    true,
		},
		schemaNamespaceKey: {
			Type:        schema.TypeString,
			Description: "Namespace name",
			Required:    true,
			ForceNew:    true,
		},

		schemaConfigYAMLKey: {
			Type:        schema.TypeString,
			Description: "Configuration as YAML",
			Optional:    true,
			Sensitive:   true,
		},
		schemaFilesKey: {
			Type:        schema.TypeList,
			Description: "Files",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		schemaDiffChangesKey: {
			Type:        schema.TypeBool,
			Description: "Show changes",
			Optional:    true,
		},
		schemaDiffContextKey: {
			Type:        schema.TypeInt,
			Description: "Show number of lines around changed lines",
			Optional:    true,
		},

		schemaDebugLogsKey: {
			Type:        schema.TypeBool,
			Description: "Enable debug logging",
			Optional:    true,
			Default:     false,
		},

		schemaDeployKey: {
			Type:        schema.TypeList,
			Description: "Deploy options",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					schemaRawOptionsKey: {
						Type:        schema.TypeList,
						Description: "Raw options given to kapp",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},

		schemaDeleteKey: {
			Type:        schema.TypeList,
			Description: "Delete options",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					schemaRawOptionsKey: {
						Type:        schema.TypeList,
						Description: "Raw options given to kapp",
						Optional:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},

		// Fields automatically managed by the kapp resource
		schemaClusterDriftDetectedKey: {
			Type:        schema.TypeBool,
			Description: "Internal (forces resource update when detected cluster drift)",
			Optional:    true,
			Default:     false,
		},
		schemaDiffPreview1Key: { // Used during create operations
			Type:        schema.TypeString,
			Description: "Shows calculated diff (1)",
			Computed:    true,
		},
		schemaDiffPreview2Key: { // Used during update operations
			Type:        schema.TypeString,
			Description: "Shows calculated diff (2)",
			Optional:    true,
			Default:     "",
		},
	}
)
