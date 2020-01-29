## k14s_ytt

k14s_ytt data source provides ability to template with ytt.

### Input Attributes

- `files` (list of strings) List of file paths (and/or directory paths) to provie to ytt
- `ignore_unknown_comments` (bool; optional) Equivalent to --ignore-unkown-comments
- `values` (map[string]string) Data values as _string_ KVs. Passed to ytt via `--data-value` flag. Note that terraform allows to set values to any type (e.g. boolean); however, it will coerce values to strings before this provider seem them.
- `values_yaml` (map[string]string) Data values as YAML KVs. Passed to ytt via `--data-value-yaml`.
- `config_yaml` (string) Configuration YAML (multiline strings are indent-trimmed). Could include YAML document annotated as `@data/values`. Passed to ytt over stdin.
- `debug_logs` (bool; optional; default=false) Log to /tmp/terraform-provider-k14s.log

### Computed Attributes

- `result` (string) Output of ytt templating operation

### Example

Use `values` for string values and `values_yaml` for non-string values.

```yaml
data "k14s_ytt" "tpl1" {
  files = ["ytt-k8s"]

  values = {
    prop1 = "val1"
  }

  values_yaml = {
    prop2 = true
    prop3 = 156
  }
}

output "result" {
  value = "${data.k14s_ytt.tpl1.result}"
}
```

Use `config_yaml` to provide more complex data values

```yaml
data "k14s_ytt" "tpl1" {
  files = ["ytt-k8s"]

  config_yaml = <<EOF
    #@data/values
    ---
    cm1: "cm1"
    cm2: "cm3"
  EOF

  values = {
    cm1 = "cm1-updated"
    "cm3.nested" = "cm3-updated"
  }
}

output "result" {
  value = "${data.k14s_ytt.tpl1.result}"
}
```
