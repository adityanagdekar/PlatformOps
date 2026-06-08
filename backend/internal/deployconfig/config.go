package deployconfig

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type DeploymentConfig struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Image     string            `yaml:"image"`
	Tag       string            `yaml:"tag"`
	Replicas  int32             `yaml:"replicas"`
	Ports     PortConfig        `yaml:"ports"`
	Health    HealthConfig      `yaml:"health"`
	Resources ResourceConfig    `yaml:"resources"`
	Env       map[string]string `yaml:"env"`
}

type PortConfig struct {
	Container int32 `yaml:"container"`
	Service   int32 `yaml:"service"`
}

type HealthConfig struct {
	Path                string `yaml:"path"`
	InitialDelaySeconds int32  `yaml:"initialDelaySeconds"`
	PeriodSeconds       int32  `yaml:"periodSeconds"`
}

type ResourceConfig struct {
	CPURequest    string `yaml:"cpuRequest"`
	MemoryRequest string `yaml:"memoryRequest"`
	CPULimit      string `yaml:"cpuLimit"`
	MemoryLimit   string `yaml:"memoryLimit"`
}

func Load(path string) (*DeploymentConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read deployment config: %w", err)
	}

	var config DeploymentConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parse deployment config: %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c DeploymentConfig) FullImage() string {
	image := strings.TrimSpace(c.Image)
	tag := strings.TrimSpace(c.Tag)
	if tag == "" || strings.Contains(image, ":") {
		return image
	}

	return image + ":" + tag
}

func (c DeploymentConfig) Validate() error {
	var missing []string

	if strings.TrimSpace(c.Name) == "" {
		missing = append(missing, "name")
	}
	if strings.TrimSpace(c.Namespace) == "" {
		missing = append(missing, "namespace")
	}
	if strings.TrimSpace(c.Image) == "" {
		missing = append(missing, "image")
	}
	if c.Replicas < 1 {
		missing = append(missing, "replicas")
	}
	if c.Ports.Container < 1 {
		missing = append(missing, "ports.container")
	}
	if c.Ports.Service < 1 {
		missing = append(missing, "ports.service")
	}
	if strings.TrimSpace(c.Health.Path) == "" {
		missing = append(missing, "health.path")
	}

	if len(missing) > 0 {
		return fmt.Errorf("invalid deployment config: missing or invalid %s", strings.Join(missing, ", "))
	}

	return nil
}

func IsNotFound(err error) bool {
	return errors.Is(err, os.ErrNotExist)
}
