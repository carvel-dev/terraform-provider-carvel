#!/bin/bash

set -e -x -u

# makes builds reproducible
export CGO_ENABLED=0
repro_flags="-ldflags=-buildid= -trimpath -mod=vendor"

go fmt ./cmd/... ./pkg/...
go mod vendor
go mod tidy

go build $repro_flags -o terraform-provider-k14s ./cmd/...
ls -la ./terraform-provider-k14s

echo SUCCESS
