#!/bin/sh
set -eu

if ! command -v go >/dev/null 2>&1; then
  printf 'go is required but was not found on PATH.\n'
  printf 'Install Go 1.22+, then rerun: make backend-run\n'
  exit 1
fi

cd backend
go run ./cmd/server

