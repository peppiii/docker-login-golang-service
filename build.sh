#!/bin/bash
set -e
set -o pipefail
## go clean && CGO_ENABLED=0 go build
env GOOS=linux GOARCH=386 go build
docker build -t docker-login-service:v1.0.0 .
