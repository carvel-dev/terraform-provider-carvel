## Install

As of v0.10.0+, Carvel provider is published to Terraform Registry: https://registry.terraform.io/providers/vmware-tanzu/carvel/latest

Depending on the resource you are planning to use (ytt, kbld, kapp), you will have to install those projects' binaries and make them available on $PATH. See [carvel.dev](https://carvel.dev) for installation instructions.

If you would like to use this provider on Terraform Cloud, recommended approach (at this point) is to include used binaries (ytt, kbld, kapp...) in a separate Git repository as a submodule to your project and configure $PATH to pick them up.
