#!/bin/sh
set -eu

CLUSTER_NAME="${CLUSTER_NAME:-platformops}"
K8S_DIR="${K8S_DIR:-examples/sample-api/k8s}"

require_tool() {
  if ! command -v "$1" >/dev/null 2>&1; then
    printf '%s is required but was not found on PATH.\n' "$1"
    exit 1
  fi
}

require_tool kind
require_tool kubectl

if ! kind get clusters | grep -qx "$CLUSTER_NAME"; then
  printf 'Kind cluster does not exist: %s\n' "$CLUSTER_NAME"
  exit 1
fi

kubectl delete -f "$K8S_DIR/service.yaml" --ignore-not-found
kubectl delete -f "$K8S_DIR/deployment.yaml" --ignore-not-found
kubectl delete -f "$K8S_DIR/configmap.yaml" --ignore-not-found
