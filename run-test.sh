#!/usr/bin/env bash
gotestsum --rerun-fails --packages=. -- -coverprofile=coverage.out -coverpkg="."
go tool cover -func coverage.out
