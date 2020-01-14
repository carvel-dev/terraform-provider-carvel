## k14s_ytt

k14s_ytt data source provides ability to template with ytt.

### Input Attributes

- `files` (list of strings) List of file paths (and/or directory paths) to provie to ytt
- `ignore_unknown_comments` (bool; optional) Equivalent to --ignore-unkown-comments
- `values_yaml` (string) Data values as YAML (multiline strings are indent-trimmed)
- `debug_logs` (bool; optional; default=false) Log to /tmp/terraform-provider-k14s.log

### Computed Attributes

- `result` (string) Output of ytt templating operation

### Example

```yaml
data "k14s_ytt" "tpl1" {
  files = ["ytt-k8s"]
  values_yaml = <<EOF
    #@data/values
    ---
    cm1: "cm1"
    cm2: "cm3"
  EOF
}

output "result" {
  value = "${data.k14s_ytt.tpl1.result}"
}
```
