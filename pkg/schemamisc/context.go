package schemamisc

type Kubeconfig interface {
	AsString() (string, string, error)
}

type Context struct {
	Kubeconfig  Kubeconfig
	DiffPreview bool
}
