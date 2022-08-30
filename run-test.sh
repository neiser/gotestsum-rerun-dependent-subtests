#!/usr/bin/env bash
go install gotest.tools/gotestsum@rerun-root-cases
gotestsum --debug --rerun-fails --rerun-fails-run-root-test --packages=. -- -coverprofile=coverage.out -coverpkg="."
go tool cover -func coverage.out
