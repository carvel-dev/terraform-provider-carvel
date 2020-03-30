## Kubernetes Guestbook example

Files:

- `app.tf`: specifies Terraform config
- `*.yaml`: specifies guestbook application (configuration copied from https://github.com/kubernetes/examples/tree/f3d89d074fe992d12adb54ad9859a68fe1e1e082/guestbook/all-in-one)
- `terraform-provider-k14s`: shim script for terraform provider (only necessary for quick provider iteration)

Example output from running `terraform apply`:

```bash
$ git clone https://github.com/k14s/terraform-provider-k14s

$ cd terraform-provider-k14s/examples/guestbook

$ terraform apply
data.k14s_ytt.guestbook: Refreshing state...

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # k14s_kapp.guestbook will be created
  + resource "k14s_kapp" "guestbook" {
      + app                    = "guestbook"
      + change_diff            = <<~EOT
            Target cluster 'https://35.239.160.185' (nodes: gke-dk-jan-9-default-pool-a218b1c9-55sl, 4+)

            --- create service/frontend (v1) namespace: default
                  0 + apiVersion: v1
                  1 + kind: Service
                  2 + metadata:
                  3 +   labels:
                  4 +     app: guestbook
                  5 +     kapp.k14s.io/app: "1580317584170216000"
                  6 +     kapp.k14s.io/association: v1.9f8d036156bc7b497df91a6a7d2b48cd
                  7 +     tier: frontend
                  8 +   name: frontend
                  9 +   namespace: default
                 10 + spec:
                 11 +   ports:
                 12 +   - port: 80
                 13 +   selector:
                 14 +     app: guestbook
                 15 +     kapp.k14s.io/app: "1580317584170216000"
                 16 +     tier: frontend
                 17 +
            --- create deployment/frontend (apps/v1) namespace: default
                  0 + apiVersion: apps/v1
                  1 + kind: Deployment
                  2 + metadata:
                  3 +   labels:
                  4 +     kapp.k14s.io/app: "1580317584170216000"
                  5 +     kapp.k14s.io/association: v1.95c1511bde234f3b1296c5e2be3c6864
                  6 +   name: frontend
                  7 +   namespace: default
                  8 + spec:
                  9 +   replicas: 1
                 10 +   selector:
                 11 +     matchLabels:
                 12 +       app: guestbook
                 13 +       kapp.k14s.io/app: "1580317584170216000"
                 14 +       tier: frontend
                 15 +   template:
                 16 +     metadata:
                 17 +       labels:
                 18 +         app: guestbook
                 19 +         kapp.k14s.io/app: "1580317584170216000"
                 20 +         kapp.k14s.io/association: v1.95c1511bde234f3b1296c5e2be3c6864
                 21 +         tier: frontend
                 22 +     spec:
                 23 +       containers:
                 24 +       - env:
                 25 +         - name: GET_HOSTS_FROM
                 26 +           value: dns
                 27 +         image: gcr.io/google-samples/gb-frontend:v4
                 28 +         name: php-redis
                 29 +         ports:
                 30 +         - containerPort: 80
                 31 +         resources:
                 32 +           requests:
                 33 +             cpu: 100m
                 34 +             memory: 100Mi
                 35 +
            --- create service/redis-master (v1) namespace: default
                  0 + apiVersion: v1
                  1 + kind: Service
                  2 + metadata:
                  3 +   labels:
                  4 +     app: redis
                  5 +     kapp.k14s.io/app: "1580317584170216000"
                  6 +     kapp.k14s.io/association: v1.b48f98222b539561040c68ac84cb5ba1
                  7 +     role: master
                  8 +     tier: backend
                  9 +   name: redis-master
                 10 +   namespace: default
                 11 + spec:
                 12 +   ports:
                 13 +   - port: 6379
                 14 +     targetPort: 6379
                 15 +   selector:
                 16 +     app: redis
                 17 +     kapp.k14s.io/app: "1580317584170216000"
                 18 +     role: master
                 19 +     tier: backend
                 20 +
            --- create deployment/redis-master (apps/v1) namespace: default
                  0 + apiVersion: apps/v1
                  1 + kind: Deployment
                  2 + metadata:
                  3 +   labels:
                  4 +     kapp.k14s.io/app: "1580317584170216000"
                  5 +     kapp.k14s.io/association: v1.dacdd80c0159ff01f8bdbbcd922ed2bd
                  6 +   name: redis-master
                  7 +   namespace: default
                  8 + spec:
                  9 +   replicas: 1
                 10 +   selector:
                 11 +     matchLabels:
                 12 +       app: redis
                 13 +       kapp.k14s.io/app: "1580317584170216000"
                 14 +       role: master
                 15 +       tier: backend
                 16 +   template:
                 17 +     metadata:
                 18 +       labels:
                 19 +         app: redis
                 20 +         kapp.k14s.io/app: "1580317584170216000"
                 21 +         kapp.k14s.io/association: v1.dacdd80c0159ff01f8bdbbcd922ed2bd
                 22 +         role: master
                 23 +         tier: backend
                 24 +     spec:
                 25 +       containers:
                 26 +       - image: k8s.gcr.io/redis:e2e
                 27 +         name: master
                 28 +         ports:
                 29 +         - containerPort: 6379
                 30 +         resources:
                 31 +           requests:
                 32 +             cpu: 100m
                 33 +             memory: 100Mi
                 34 +
            --- create service/redis-slave (v1) namespace: default
                  0 + apiVersion: v1
                  1 + kind: Service
                  2 + metadata:
                  3 +   labels:
                  4 +     app: redis
                  5 +     kapp.k14s.io/app: "1580317584170216000"
                  6 +     kapp.k14s.io/association: v1.fb72c55fcff353bdbf28c0d9a75196d6
                  7 +     role: slave
                  8 +     tier: backend
                  9 +   name: redis-slave
                 10 +   namespace: default
                 11 + spec:
                 12 +   ports:
                 13 +   - port: 6379
                 14 +   selector:
                 15 +     app: redis
                 16 +     kapp.k14s.io/app: "1580317584170216000"
                 17 +     role: slave
                 18 +     tier: backend
                 19 +
            --- create deployment/redis-slave (apps/v1) namespace: default
                  0 + apiVersion: apps/v1
                  1 + kind: Deployment
                  2 + metadata:
                  3 +   labels:
                  4 +     kapp.k14s.io/app: "1580317584170216000"
                  5 +     kapp.k14s.io/association: v1.4a90a0c0db348752b3fc21e2271f78d4
                  6 +   name: redis-slave
                  7 +   namespace: default
                  8 + spec:
                  9 +   replicas: 1
                 10 +   selector:
                 11 +     matchLabels:
                 12 +       app: redis
                 13 +       kapp.k14s.io/app: "1580317584170216000"
                 14 +       role: slave
                 15 +       tier: backend
                 16 +   template:
                 17 +     metadata:
                 18 +       labels:
                 19 +         app: redis
                 20 +         kapp.k14s.io/app: "1580317584170216000"
                 21 +         kapp.k14s.io/association: v1.4a90a0c0db348752b3fc21e2271f78d4
                 22 +         role: slave
                 23 +         tier: backend
                 24 +     spec:
                 25 +       containers:
                 26 +       - env:
                 27 +         - name: GET_HOSTS_FROM
                 28 +           value: dns
                 29 +         image: gcr.io/google_samples/gb-redisslave:v1
                 30 +         name: slave
                 31 +         ports:
                 32 +         - containerPort: 6379
                 33 +         resources:
                 34 +           requests:
                 35 +             cpu: 100m
                 36 +             memory: 100Mi
                 37 +

            Changes

            Namespace  Name          Kind        Conds.  Age  Op      Wait to    Rs  Ri
            default    frontend      Deployment  -       -    create  reconcile  -   -
            ^          frontend      Service     -       -    create  reconcile  -   -
            ^          redis-master  Deployment  -       -    create  reconcile  -   -
            ^          redis-master  Service     -       -    create  reconcile  -   -
            ^          redis-slave   Deployment  -       -    create  reconcile  -   -
            ^          redis-slave   Service     -       -    create  reconcile  -   -

            Op:      6 create, 0 delete, 0 update, 0 noop
            Wait to: 6 reconcile, 0 delete, 0 noop

        EOT
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
