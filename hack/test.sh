#!/bin/bash

set -e -x -u

./hack/build.sh

cd examples/

terraform init
terraform apply -auto-approve
terraform destroy -auto-approve
