package kapp

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/logger"
)

type SettableDiff struct {
	diff   *schema.ResourceDiff
	logger logger.Logger
}

func (d SettableDiff) Get(key string) interface{} {
	return d.diff.Get(key)
}

func (d SettableDiff) GetOk(key string) (interface{}, bool) {
	return d.diff.GetOk(key)
}

func (d SettableDiff) Set(key string, val interface{}) error {
	if keySchema, found := resourceSchema[key]; found {
		if keySchema.Computed {
			return d.diff.SetNew(key, val)
		} else {
			d.logger.Debug(fmt.Sprintf("Skip setting %s (not computed)", key))
		}
		return nil
	}
	panic(fmt.Sprintf("Expected to find schema for key '%s'", key))
}
