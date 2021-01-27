## Install

[Terraform docs](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) for installing third party providers.

Grab prebuilt binaries from the [Releases page](https://github.com/vmware-tanzu/terraform-provider-carvel/releases).

Once you have downloaded `terraform-provider-carvel-binaries.tgz`, install it for terraform to find it:

```bash
mkdir -p ~/.terraform.d/plugins
tar xzvf ~/Downloads/terraform-provider-carvel-binaries.tgz -C ~/.terraform.d/plugins/
```

Depending on the resource you are planning to use (ytt, kbld, kapp), you will have to install those projects' binaries. See [carvel.dev](https://carvel.dev) for installation instructions.
