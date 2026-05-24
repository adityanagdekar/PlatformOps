SHELL := /bin/sh

CLUSTER_NAME ?= platformops
BACKEND_IMAGE ?= platformops-backend:local

.PHONY: check-tools
check-tools:
	./scripts/check-tools.sh

.PHONY: cluster-up
cluster-up:
	./scripts/cluster-up.sh

.PHONY: cluster-down
cluster-down:
	./scripts/cluster-down.sh

.PHONY: backend-run
backend-run:
	./scripts/backend-run.sh

.PHONY: backend-test
backend-test:
	cd backend && go test ./...

.PHONY: backend-docker-build
backend-docker-build:
	docker build -t $(BACKEND_IMAGE) ./backend

.PHONY: deploy-sample
deploy-sample:
	./scripts/deploy-sample.sh

.PHONY: monitoring-up
monitoring-up:
	./scripts/monitoring-up.sh
