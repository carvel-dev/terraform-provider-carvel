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
