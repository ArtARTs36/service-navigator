package config

import (
	"fmt"
	"log"
	"os"
	"time"

	app "github.com/artarts36/service-navigator/internal/application"
	"github.com/artarts36/service-navigator/internal/presentation/config"
	"gopkg.in/yaml.v3"
)

const serviceMetricDepth = 50
const servicePollInterval = 1 * time.Second

type Config struct {
	Frontend config.Frontend `yaml:"frontend"`
	Backend  struct {
		NetworkName string           `yaml:"network_name"`
		Poll        app.PollerConfig `yaml:"poll"`
		Images      struct {
			Poll app.ImagePollerConfig
		} `yaml:"images"`
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

	conf.Backend.Poll.Metrics.Depth = resolveConfigMetricDepth(conf.Backend.Poll)

	if conf.Backend.Poll.Interval == 0 || conf.Backend.Poll.Interval < 0 {
		conf.Backend.Poll.Interval = servicePollInterval
	}

	if conf.Frontend.AppName == "" {
		conf.Frontend.AppName = "ServiceNavigator"
	}

	log.Printf("Config loaded: %v", conf)

	return &conf
}

func resolveConfigMetricDepth(conf app.PollerConfig) int {
	if conf.Metrics.Depth > 0 {
		return conf.Metrics.Depth
	}

	if conf.Metrics.Depth == 0 {
		return serviceMetricDepth
	}

	return 1
}
