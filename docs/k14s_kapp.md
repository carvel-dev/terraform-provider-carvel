## k14s_kapp

k14s_kapp resource provides ability to manage set of Kubernetes resources.

### Input Attributes

- `app` (string; required) App name
- `namespace` (string; required) Namespace name
- `config_yaml` (string; optional; sensitive) Input configuration as YAML (multiline strings are indent-trimmed)
- `files` (list of strings; optional) List of file paths to provide to kapp
- `diff_changes` (bool; optional) Equivalent to --diff-changes
- `diff_context` (int; optional) Equivalent to --diff-context
- `debug_logs` (bool; optional; default=false) Log to /tmp/terraform-provider-k14s.log

### Computed Attributes

- `cluster_drift_detected` (bool) Set to true when kapp detects there are non-matching changes in the cluster compared to provided configuration
- `change_diff` (string) Shows diff output from kapp

### Example

```yaml
data "k14s_kapp" "app2" {
  app = "app2"
  namespace = "default"

  config_yaml = <<EOF
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: app2
    data:
      data.txt: something
  EOF

  diff_changes = true
}
```
