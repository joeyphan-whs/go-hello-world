#!/usr/bin/env bash

set -e


image() {
    docker build -t go-hello-world-build:latest -f build/docker/Dockerfile.build .
    docker build -t hello-world:latest -f build/docker/Dockerfile .
}

test() {
    ginkgo -cover ./cmd/go-hello-world
}

: ${1?}
$@
