#!/bin/sh
set -eu

CLUSTER_NAME="${CLUSTER_NAME:-platformops}"

if ! command -v kind >/dev/null 2>&1; then
  printf 'kind is required but was not found on PATH.\n'
  exit 1
fi

if kind get clusters | grep -qx "$CLUSTER_NAME"; then
  kind delete cluster --name "$CLUSTER_NAME"
else
  printf 'Kind cluster does not exist: %s\n' "$CLUSTER_NAME"
fi

