package schemamisc

import (
	"github.com/vmware-tanzu/terraform-provider-carvel/pkg/logger"
)

type Kubeconfig interface {
	AsString() (string, string, error)
}

type Context struct {
	Kubeconfig        Kubeconfig
	DiffPreviewLogger logger.Logger
}
