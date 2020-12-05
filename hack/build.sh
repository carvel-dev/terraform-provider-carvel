#!/bin/bash

set -e -x -u

go fmt ./cmd/... ./pkg/...
go mod vendor
go mod tidy
go build -o terraform-provider-k14s ./cmd/... 
ls -la ./terraform-provider-k14s

echo SUCCESS
