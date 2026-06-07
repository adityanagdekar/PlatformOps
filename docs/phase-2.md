# Phase 2: Sample App Deployment

## Objective

Create one simple backend service that can be built as a Docker image and manually deployed to the local Kind cluster with Kubernetes manifests.

## Included

- Sample Go API under `examples/sample-api`
- `GET /health`
- `GET /ready`
- `GET /metrics`
- Sample API Dockerfile
- Kubernetes Namespace, ConfigMap, Deployment, and Service
- readinessProbe on `/ready`
- livenessProbe on `/health`
- Kind image load and `kubectl apply` deployment script

## Validation

Expected commands:

```sh
make backend-test
make sample-docker-build
make cluster-up
make deploy-sample
kubectl get pods -n platformops
kubectl get svc -n platformops
kubectl port-forward svc/sample-api -n platformops 8081:80
curl http://localhost:8081/health
curl http://localhost:8081/ready
curl http://localhost:8081/metrics
```

Cleanup:

```sh
make undeploy-sample
```
