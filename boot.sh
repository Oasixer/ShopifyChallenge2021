#!/bin/sh

cd "${0%/*}/backend" # make sure cwd is correct for file paths
./main --migrate-serve
