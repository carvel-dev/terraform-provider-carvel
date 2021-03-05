data "carvel_kbld" "tpl1" {
  config_yaml = <<EOF
    images:
    - image: nginx
    - image: mysql
  EOF
}

output "kbld_result" {
  value = data.carvel_kbld.tpl1.result
}
