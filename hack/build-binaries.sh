#!/bin/bash

set -e -x -u

./hack/build.sh

# makes builds reproducible
export CGO_ENABLED=0
repro_flags="-ldflags=-buildid= -trimpath"

GOOS=darwin GOARCH=amd64 go build $repro_flags -o terraform-provider-k14s-darwin-amd64 ./cmd/...
GOOS=linux GOARCH=amd64 go build $repro_flags -o terraform-provider-k14s-linux-amd64 ./cmd/...
GOOS=windows GOARCH=amd64 go build $repro_flags -o terraform-provider-k14s-windows-amd64.exe ./cmd/...

shasum -a 256 ./terraform-provider-k14s-*-amd64*
