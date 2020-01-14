resource "k14s_kapp" "app1" {
  app = "app1"
  namespace = "default"

  config_yaml = <<EOF
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: str-123
    data:
      str: str-124
  EOF

  diff_changes = true
}
