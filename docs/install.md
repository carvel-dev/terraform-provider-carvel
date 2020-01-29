## Install

[Terraform docs](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) for installing third party providers.

Grab prebuilt binaries from the [Releases page](https://github.com/k14s/terraform-provider-k14s/releases).

Once you have downloaded `terraform-provider-k14s-binaries.tgz`, install it for terraform to find it:

```bash
mkdir -p ~/.terraform.d/plugins
tar xzvf ~/Downloads/terraform-provider-k14s-binaries.tgz -C ~/.terraform.d/plugins/
```

Depending on the resource you are planning to use (ytt, kbld, kapp), you will have to install those projects' binaries. See [k14s.io](https://k14s.io) for installation instructions.
