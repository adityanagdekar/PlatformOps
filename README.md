# PlatformOps

PlatformOps is a lightweight internal developer platform for deploying, observing, and managing containerized backend services on a local Kubernetes cluster.

Currently the project has a runnable Go backend skeleton, local automation for a Kind-based Kubernetes environment, and a manually deployable sample API.

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

Implemented in Phase 2:

- Sample Go backend service with `GET /health`, `GET /ready`, and `GET /metrics`
- Dockerfile for the sample service
- Kubernetes Namespace, ConfigMap, Deployment, and Service manifests
- Readiness and liveness probes
- Local Kind deployment script for the sample service

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
make sample-docker-build
make deploy-sample
make undeploy-sample
make cluster-down
```

`make monitoring-up` exists as a placeholder for the later monitoring phase.

## Backend

Run locally:

```sh
make backend-run
```

Health check:

```sh
curl http://localhost:8080/health
```

## Sample API

Build the sample image:

```sh
make sample-docker-build
```

Deploy it to the local Kind cluster:

```sh
make cluster-up
make deploy-sample
```

Inspect Kubernetes resources:

```sh
kubectl get pods -n platformops
kubectl get svc -n platformops
```

Port-forward the sample service:

```sh
kubectl port-forward svc/sample-api -n platformops 8081:80
```

Verify the service from another terminal:

```sh
curl http://localhost:8081/health
curl http://localhost:8081/ready
curl http://localhost:8081/metrics
```

Remove sample app resources:

```sh
make undeploy-sample
```
