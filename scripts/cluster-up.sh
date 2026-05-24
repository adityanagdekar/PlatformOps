#!/bin/sh
set -eu

CLUSTER_NAME="${CLUSTER_NAME:-platformops}"
CONFIG_FILE="${CONFIG_FILE:-k8s/kind-config.yaml}"

if ! command -v kind >/dev/null 2>&1; then
  printf 'kind is required but was not found on PATH.\n'
  printf 'Install Kind, then rerun: make cluster-up\n'
  exit 1
fi

if ! command -v kubectl >/dev/null 2>&1; then
  printf 'kubectl is required but was not found on PATH.\n'
  exit 1
fi

if kind get clusters | grep -qx "$CLUSTER_NAME"; then
  printf 'Kind cluster already exists: %s\n' "$CLUSTER_NAME"
else
  kind create cluster --name "$CLUSTER_NAME" --config "$CONFIG_FILE"
fi

kubectl cluster-info --context "kind-$CLUSTER_NAME"

