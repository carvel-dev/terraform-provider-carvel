#!/bin/bash

set -e -x -u

go fmt ./cmd/... ./pkg/...

go build -o terraform-provider-k14s ./cmd/...
ls -la ./terraform-provider-k14s

echo SUCCESS
