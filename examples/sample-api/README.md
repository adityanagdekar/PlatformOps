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

Container build:

```sh
docker build -t sample-api:local ./examples/sample-api
```

Deploy to Kind from the repository root:

```sh
make cluster-up
make deploy-sample
```

Inspect the deployment:

```sh
kubectl get pods -n platformops
kubectl get svc -n platformops
kubectl rollout status deployment/sample-api -n platformops
```

Port-forward and verify:

```sh
kubectl port-forward svc/sample-api -n platformops 8081:80
curl http://localhost:8081/health
curl http://localhost:8081/ready
curl http://localhost:8081/metrics
```
