#!/bin/bash

set -e -x -u

export VERSION=0.7.0

./hack/build.sh

rm -rf tmp/binaries/k14s
mkdir -p tmp/binaries/k14s/$VERSION

(
	set -e

	cd tmp/binaries/k14s/$VERSION
	
	mkdir {darwin_amd64,linux_amd64,windows_amd64}

	# makes builds reproducible
	export CGO_ENABLED=0
	repro_flags="-ldflags=-buildid= -trimpath"

	GOOS=darwin GOARCH=amd64 go build $repro_flags \
		-o darwin_amd64/terraform-provider-k14s ../../../../cmd/...
	GOOS=linux GOARCH=amd64 go build $repro_flags \
		-o linux_amd64/terraform-provider-k14s ../../../../cmd/...
	GOOS=windows GOARCH=amd64 go build $repro_flags \
		-o windows_amd64/terraform-provider-k14s ../../../../cmd/...

	cd ../../
	COPYFILE_DISABLE=1 tar czvf ../terraform-provider-k14s-binaries.tgz .
)

shasum -a 256 tmp/terraform-provider-k14s-binaries.tgz
