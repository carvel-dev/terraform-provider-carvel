![logo](docs/CarvelLogo.png)

# terraform-provider-k14s

- Slack: [#carvel in Kubernetes slack](https://slack.kubernetes.io)
- [Docs](docs/README.md) with topics about resources and their attributes, examples
- Refer to [Install doc](docs/install.md) for setup instructions
- Status: Experimental

k14s Terraform provider currently includes ability to:

- template with [ytt](https://get-ytt.io)
- resolve image digests with [kbld](https://get-kbld.io)
- deploy k8s resources with [kapp](https://get-kapp.io)

See [examples/guestbook](examples/guestbook) for an example installing Kubernetes Guestbook.

Works with Terraform 0.13+

### Join the Community and Make Carvel Better
Carvel is better because of our contributors and maintainers. It is because of you that we can bring great software to the community.
Please join us during our online community meetings. Details can be found on our [Carvel website](https://carvel.dev/community/).

You can chat with us on Kubernetes Slack in the #carvel channel and follow us on Twitter at @carvel_dev.

Check out which organizations are using and contributing to Carvel: [Adopter's list](https://github.com/vmware-tanzu/carvel/blob/master/ADOPTERS.md)
