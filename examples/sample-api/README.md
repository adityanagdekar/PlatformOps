# sample-api

This directory contains the Phase 2 deployable backend service.

Endpoints:

- `GET /health`
- `GET /ready`
- `GET /metrics`

Run locally:

```sh
cd examples/sample-api
go run .
```

Then verify:

```sh
curl http://localhost:8080/health
curl http://localhost:8080/ready
curl http://localhost:8080/metrics
```

Planned runtime assets for the rest of Phase 2:

- `Dockerfile`
- Kubernetes `Deployment`
- Kubernetes `Service`
- Kubernetes `ConfigMap`
- readiness and liveness probes
