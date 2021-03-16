provider "carvel" {
  kapp {
    kubeconfig {
      from_env = true
    }
  }
}

data "carvel_ytt" "guestbook" {
  files = ["ytt-config"]
  ignore_unknown_comments = true

  # Configure all deployments to have 1 replica
  config_yaml = <<EOF
    #@ load("@ytt:overlay", "overlay")
    #@overlay/match by=overlay.subset({"kind":"Deployment"}),expects="1+"
    ---
    spec:
      replicas: 1
  EOF
}

resource "carvel_kapp" "guestbook" {
  app = "guestbook"
  namespace = "default"
  config_yaml = data.carvel_ytt.guestbook.result
  diff_changes = true
}
