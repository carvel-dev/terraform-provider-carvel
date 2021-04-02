#!/bin/bash

set -e -x -u

export VERSION=0.8.0 # No v prefix

./hack/build.sh

rm -rf tmp/binaries/carvel
mkdir -p tmp/binaries/carvel/$VERSION

(
	set -e

	cd tmp/binaries/carvel/$VERSION
	
	mkdir {darwin_amd64,linux_amd64,windows_amd64}

	# makes builds reproducible
	export CGO_ENABLED=0
	repro_flags="-ldflags=-buildid= -trimpath -mod=vendor"

	GOOS=darwin GOARCH=amd64 go build $repro_flags \
		-o darwin_amd64/terraform-provider-carvel ../../../../cmd/...
	GOOS=linux GOARCH=amd64 go build $repro_flags \
		-o linux_amd64/terraform-provider-carvel ../../../../cmd/...
	GOOS=windows GOARCH=amd64 go build $repro_flags \
		-o windows_amd64/terraform-provider-carvel ../../../../cmd/...

	cd ../../
	COPYFILE_DISABLE=1 tar czvf ../terraform-provider-carvel-binaries.tgz .
)

shasum -a 256 tmp/terraform-provider-carvel-binaries.tgz
