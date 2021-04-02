## Kubernetes Guestbook example

Files:

- `app.tf`: specifies Terraform config
- `*.yaml`: specifies guestbook application (configuration copied from https://github.com/kubernetes/examples/tree/f3d89d074fe992d12adb54ad9859a68fe1e1e082/guestbook/all-in-one)
- `terraform-provider-carvel`: shim script for terraform provider (only necessary for quick provider iteration)

Example output from running `terraform apply`:

```bash
$ git clone https://github.com/vmware-tanzu/terraform-provider-carvel

$ cd terraform-provider-carvel/examples/guestbook

$ terraform apply
data.carvel_ytt.guestbook: Refreshing state...

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # carvel_kapp.guestbook will be created
  + resource "carvel_kapp" "guestbook" {
      + app                    = "guestbook"
      + cluster_drift_detected = false
      + config_yaml            = (sensitive value)
      + debug_logs             = false
      + diff_changes           = true
      + id                     = (known after apply)
      + namespace              = "default"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value:
```
