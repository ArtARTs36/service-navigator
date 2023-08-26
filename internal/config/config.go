package config

import (
	"fmt"
	"log"
	"os"
	"time"

	poller "github.com/artarts36/service-navigator/internal/application"
	"github.com/artarts36/service-navigator/internal/presentation/config"
	"gopkg.in/yaml.v3"
)

const serviceMetricDepth = 50
const servicePollInterval = 1 * time.Second

type Config struct {
	Frontend config.Frontend `yaml:"frontend"`
	Backend  struct {
		NetworkName string              `yaml:"network_name"`
		Poll        poller.PollerConfig `yaml:"poll"`
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

	conf.Backend.Poll.Metrics.Depth = resolveConfigMetricDepth(conf.Backend.Poll)

	if conf.Backend.Poll.Interval == 0 || conf.Backend.Poll.Interval < 0 {
		conf.Backend.Poll.Interval = servicePollInterval
	}

	return &conf
}

func resolveConfigMetricDepth(conf poller.PollerConfig) int {
	if conf.Metrics.Depth > 0 {
		return conf.Metrics.Depth
	}

	if conf.Metrics.Depth == 0 {
		return serviceMetricDepth
	}

	return 1
}
