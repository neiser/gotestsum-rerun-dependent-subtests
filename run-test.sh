#!/usr/bin/env bash
set -eu
go install gotest.tools/gotestsum@rerun-root-cases
go install github.com/neiser/gocovmerge@latest
rm coverage.out.rerun.*
gotestsum --rerun-fails --rerun-fails-run-root-test --packages=. --raw-command -- ./test-with-coverage.sh
gocovmerge coverage.out.rerun.* > coverage.out
go tool cover -func coverage.out
