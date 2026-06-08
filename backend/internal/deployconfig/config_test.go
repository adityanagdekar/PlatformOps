package deployconfig

import (
	"path/filepath"
	"testing"
)

func TestLoadSampleAPIConfig(t *testing.T) {
	config, err := Load(filepath.Join("..", "..", "..", "configs", "sample-api.yaml"))
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if config.Name != "sample-api" {
		t.Fatalf("expected name sample-api, got %q", config.Name)
	}
	if config.Namespace != "platformops" {
		t.Fatalf("expected namespace platformops, got %q", config.Namespace)
	}
	if config.FullImage() != "sample-api:local" {
		t.Fatalf("expected full image sample-api:local, got %q", config.FullImage())
	}
	if config.Replicas != 2 {
		t.Fatalf("expected replicas 2, got %d", config.Replicas)
	}
	if config.Ports.Container != 8080 {
		t.Fatalf("expected container port 8080, got %d", config.Ports.Container)
	}
	if config.Ports.Service != 80 {
		t.Fatalf("expected service port 80, got %d", config.Ports.Service)
	}
	if config.Health.Path != "/health" {
		t.Fatalf("expected health path /health, got %q", config.Health.Path)
	}
	if config.Env["APP_ENV"] != "local" {
		t.Fatalf("expected APP_ENV local, got %q", config.Env["APP_ENV"])
	}
}

func TestValidateRejectsMissingRequiredFields(t *testing.T) {
	config := DeploymentConfig{}

	if err := config.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}
