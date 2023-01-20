#!/bin/bash

# ENV
export GOCACHE=/root/.cache/go-build
export GOPATH=/root/go
export GOMODCACHE=/root/go/pkg/mod
alias go="/snap/bin/go"

## Run the application
cd $(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd) && go mod download && go mod tidy && go run main.go
