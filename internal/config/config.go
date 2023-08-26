package config

import (
	"fmt"
	"log"
	"os"

	poller "github.com/artarts36/service-navigator/internal/application"
	"github.com/artarts36/service-navigator/internal/presentation/config"
	"gopkg.in/yaml.v3"
)

const serviceMetricDepth = 50

type Config struct {
	Frontend config.Frontend `yaml:"frontend"`
	Backend  struct {
		NetworkName string         `yaml:"network_name"`
		Metrics     poller.Metrics `yaml:"metrics"`
	} `yaml:"backend"`
}

func InitConfig() *Config {
	conf := Config{}

	log.Printf("Loading Config from /app/service_navigator.yaml")

	yamlContent, err := os.ReadFile("/app/service_navigator.yaml")

	if err != nil {
		panic(fmt.Sprintf("Failed to read \"/app/service_navigator.yaml\": %s", err))
	}

	err = yaml.Unmarshal(yamlContent, &conf)

	if err != nil {
		panic(fmt.Sprintf("failed to load Config: %s", err))
	}

	conf.Frontend.Navbar.Search.Providers = config.ResolveProviders(conf.Frontend.Navbar.Search.Providers)

	log.Printf("Config loaded: %v", conf)

	conf.Backend.Metrics.Depth = resolveConfigMetricDepth(conf.Backend.Metrics)

	return &conf
}

func resolveConfigMetricDepth(metrics poller.Metrics) int {
	if metrics.Depth > 0 {
		return metrics.Depth
	}

	if metrics.Depth == 0 {
		return serviceMetricDepth
	}

	return 1
}
