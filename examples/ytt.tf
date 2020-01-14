data "k14s_ytt" "example1" {
  files = ["ytt-example"]
  values_yaml = <<EOF
    #@data/values
    ---
    str: tfstr
    #@overlay/match missing_ok=True
    map:
      nested: true
  EOF
}

// See `terraform output` for the result
output "result" {
  value = "${data.k14s_ytt.example1.result}"
}
