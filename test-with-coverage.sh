#!/usr/bin/env bash
set -eu
go test -json -coverprofile="$(mktemp coverage.out.rerun.XXXXXXXXXX)" -coverpkg=. "$@"