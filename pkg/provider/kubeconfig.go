package provider

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/k14s/terraform-provider-k14s/pkg/schemamisc"
)

type Kubeconfig struct {
	data *schema.ResourceData
}

func (c Kubeconfig) AsString() (string, string, error) {
	kappRaw := c.data.Get(schemaKappKey).([]interface{})
	if len(kappRaw) == 0 {
		return "", "", fmt.Errorf("Expected non-empty provider config (key '%s')", schemaKappKey)
	}

	kapp := kappRaw[0].(map[string]interface{})

	kubeconfigYAML := kapp[schemaKappKubeconfigYAMLKey].(string)
	kubeconfigRaw := kapp[schemaKappKubeconfigKey].([]interface{})

	if len(kubeconfigYAML) > 0 {
		if len(kubeconfigRaw) != 0 {
			return "", "", fmt.Errorf("Expected empty config (key '%s')", schemaKappKubeconfigKey)
		}

		kubeconfigYAML, err := schemamisc.Heredoc{kubeconfigYAML}.StripIndent()
		if err != nil {
			return "", "", fmt.Errorf("Formatting %s: %s", schemaKappKubeconfigYAMLKey, err)
		}

		return kubeconfigYAML, "", nil
	}

	if len(kubeconfigRaw) != 1 {
		return "", "", fmt.Errorf("Expected non-empty config (key '%s')", schemaKappKubeconfigKey)
	}

	kubeconfig := kubeconfigRaw[0].(map[string]interface{})

	fromEnv := kubeconfig[schemaKappKubeconfigFromEnvKey].(bool)
	context := kubeconfig[schemaKappKubeconfigContextKey].(string)
	connInfo := KubeconfigConnInfo{
		Server:     kubeconfig[schemaKappKubeconfigServerKey].(string),
		Username:   kubeconfig[schemaKappKubeconfigUsernameKey].(string),
		Password:   kubeconfig[schemaKappKubeconfigPasswordKey].(string),
		CACert:     kubeconfig[schemaKappKubeconfigCACertKey].(string),
		ClientCert: kubeconfig[schemaKappKubeconfigClientCertKey].(string),
		ClientKey:  kubeconfig[schemaKappKubeconfigClientKeyKey].(string),
	}

	if fromEnv {
		if connInfo.HasValues() {
			return "", "", fmt.Errorf("Expected key '%s' to be the only key configured",
				schemaKappKubeconfigFromEnvKey)
		}
		return "", context, nil
	}

	if len(context) > 0 {
		return "", "", fmt.Errorf("Expected key '%s' to be configured with key '%s'",
			schemaKappKubeconfigContextKey, schemaKappKubeconfigFromEnvKey)
	}

	return c.buildConfig(connInfo)
}

func (Kubeconfig) buildConfig(connInfo KubeconfigConnInfo) (string, string, error) {
	clusterConfig := map[string]interface{}{
		"server": connInfo.Server,
	}

	if len(connInfo.CACert) > 0 {
		clusterConfig["certificate-authority-data"] = base64.StdEncoding.EncodeToString([]byte(connInfo.CACert))
	}

	userConfig := map[string]interface{}{}

	if len(connInfo.Username) > 0 {
		userConfig["username"] = connInfo.Username
	}
	if len(connInfo.Password) > 0 {
		userConfig["password"] = connInfo.Password
	}
	if len(connInfo.ClientCert) > 0 {
		userConfig["client-certificate-data"] = base64.StdEncoding.EncodeToString([]byte(connInfo.ClientCert))
	}
	if len(connInfo.ClientKey) > 0 {
		userConfig["client-key-data"] = base64.StdEncoding.EncodeToString([]byte(connInfo.ClientKey))
	}

	const (
		clusterName = "kapp-cluster"
		userName    = "kapp-user"
		contextName = "kapp-context"
	)

	mergedConfig := map[string]interface{}{
		"apiVersion":      "v1",
		"kind":            "Config",
		"current-context": contextName,
		"clusters": []interface{}{
			map[string]interface{}{
				"name":    clusterName,
				"cluster": clusterConfig,
			},
		},
		"contexts": []interface{}{
			map[string]interface{}{
				"name": contextName,
				"context": map[string]interface{}{
					"cluster": clusterName,
					"user":    userName,
				},
			},
		},
		"users": []interface{}{
			map[string]interface{}{
				"name": userName,
				"user": userConfig,
			},
		},
	}

	mergedConfigBs, err := json.Marshal(mergedConfig)
	if err != nil {
		return "", "", fmt.Errorf("Marshaling kubeconfig: %s", err)
	}

	return string(mergedConfigBs), "", nil
}

type KubeconfigConnInfo struct {
	Server     string
	Username   string
	Password   string
	CACert     string
	ClientCert string
	ClientKey  string
}

func (a KubeconfigConnInfo) HasValues() bool {
	return len(a.Server+a.Username+a.Password+a.CACert+a.ClientCert+a.ClientKey) > 0
}
