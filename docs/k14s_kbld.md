## k14s_kbld

k14s_kbld data source provides ability to resolve images references with digests.

### Input Attributes

- `files` (list of strings) List of file paths (and/or directory paths) to provie to ytt
- `config_yaml` (string) Configuration as YAML (multiline strings are indent-trimmed)
- `debug_logs` (bool; optional; default=false) Log to /tmp/terraform-provider-k14s.log

### Computed Attributes

- `result` (string) Output of kbld resolution operation

### Example

```yaml
data "k14s_kbld" "tpl1" {
  config_yaml = <<EOF
    images:
    - image: nginx
    - image: mysql
  EOF
}

output "result" {
  value = "${data.k14s_kbld.tpl1.result}"
}
```
