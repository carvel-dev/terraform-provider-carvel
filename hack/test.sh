#!/bin/bash

set -e -x -u

./hack/build.sh

cd examples/test/

terraform init
terraform apply -auto-approve
terraform destroy -auto-approve

cd ../guestbook/

terraform init
terraform apply -auto-approve
terraform destroy -auto-approve
