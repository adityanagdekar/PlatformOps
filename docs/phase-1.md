# Phase 1: Local Project Foundation

## Objective

Establish the PlatformOps repository foundation and local development workflow.

## Included

- Repository structure
- Go backend skeleton
- Backend Dockerfile
- Kind cluster configuration
- Makefile command surface
- Local scripts for tool checks and cluster lifecycle
- Placeholder sample deployment command
- Starter sample app config
- CI workflow for backend validation

## Validation

Expected commands:

```sh
make check-tools
make backend-run
make backend-test
make backend-docker-build
make deploy-sample
make cluster-up
make cluster-down
```

## Notes

The backend currently exposes only:

```txt
GET /health
```

Deployment APIs begin in the next phase.

`make deploy-sample` is intentionally a placeholder until the Phase 2 sample service and Kubernetes manifests are added.
