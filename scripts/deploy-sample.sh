#!/bin/sh
set -eu

CLUSTER_NAME="${CLUSTER_NAME:-platformops}"
SAMPLE_IMAGE="${SAMPLE_IMAGE:-sample-api:local}"
SAMPLE_DIR="${SAMPLE_DIR:-examples/sample-api}"
K8S_DIR="$SAMPLE_DIR/k8s"

require_tool() {
  if ! command -v "$1" >/dev/null 2>&1; then
    printf '%s is required but was not found on PATH.\n' "$1"
    exit 1
  fi
}

require_tool docker
require_tool kind
require_tool kubectl

if ! kind get clusters | grep -qx "$CLUSTER_NAME"; then
  printf 'Kind cluster does not exist: %s\n' "$CLUSTER_NAME"
  printf 'Start it first with: make cluster-up\n'
  exit 1
fi

docker build -t "$SAMPLE_IMAGE" "$SAMPLE_DIR"
kind load docker-image "$SAMPLE_IMAGE" --name "$CLUSTER_NAME"
kubectl apply -f "$K8S_DIR/00-namespace.yaml"
kubectl apply -f "$K8S_DIR/configmap.yaml"
kubectl apply -f "$K8S_DIR/deployment.yaml"
kubectl apply -f "$K8S_DIR/service.yaml"
kubectl rollout status deployment/sample-api -n platformops
