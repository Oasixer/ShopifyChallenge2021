#!/bin/sh

cd "${0%/*}/backend" # make sure cwd is correct for file paths

# wait for postgres db
# sleep 2

# run tests
TESTING=true CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test ./... -v
