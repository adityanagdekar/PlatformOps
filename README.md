# PlatformOps

PlatformOps is a lightweight internal developer platform for deploying, observing, and managing containerized backend services on a local Kubernetes cluster.

The Phase 1 goal is intentionally small: establish the repository structure, a runnable Go backend skeleton, and local automation for a Kind-based Kubernetes environment.

## Current Status

Implemented in Phase 1:

- Go backend skeleton with `GET /health`
- Dockerfile for the PlatformOps backend
- Kind cluster configuration
- Make targets for local workflow
- Scripted tool checks and cluster lifecycle commands
- Placeholder sample deployment command
- Starter sample deployment config
- Initial CI workflow for backend tests/build

Not implemented yet:

- Application deployment API
- Kubernetes client integration
- Prometheus/Grafana stack
- Health checker
- Rollback automation

## Repository Layout

```txt
PlatformOps/
  backend/              # Go API/control plane
  configs/              # YAML runtime/deployment configs
  docs/                 # project notes and phase docs
  examples/             # sample applications managed by PlatformOps
  k8s/                  # Kubernetes manifests and Kind config
  monitoring/           # Prometheus/Grafana assets
  scripts/              # local automation scripts
  .github/workflows/    # CI/CD workflows
```

## Prerequisites

- Docker
- kubectl
- Kind
- Go 1.22+

## Local Commands

```sh
make check-tools
make cluster-up
make backend-run
make backend-test
make backend-docker-build
make deploy-sample
make cluster-down
```

`make deploy-sample` and `make monitoring-up` exist as placeholders for later phases.

## Backend

Run locally:

```sh
make backend-run
```

Health check:

```sh
curl http://localhost:8080/health
```
