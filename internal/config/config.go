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
const servicePollInterval = 2 * time.Second
const imagePollInterval = 1 * time.Minute

type Config struct {
	Frontend config.Frontend `yaml:"frontend"`
	Backend  struct {
		NetworkName string `yaml:"network_name"`
		Services    struct {
			Poll app.ServicePollerConfig `yaml:"poll"`
		} `yaml:"services"`
		Images struct {
			Poll app.ImagePollerConfig `yaml:"poll"`
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

	conf.Backend.Services.Poll.Metrics.Depth = resolveConfigMetricDepth(conf.Backend.Services.Poll)

	if conf.Backend.Services.Poll.Interval == 0 || conf.Backend.Services.Poll.Interval < 0 {
		conf.Backend.Services.Poll.Interval = servicePollInterval
	}
	if conf.Backend.Images.Poll.Interval == 0 || conf.Backend.Images.Poll.Interval < 0 {
		conf.Backend.Images.Poll.Interval = imagePollInterval
	}

	if conf.Frontend.AppName == "" {
		conf.Frontend.AppName = "ServiceNavigator"
	}

	log.Printf("Config loaded: %v", conf)

	return &conf
}

func resolveConfigMetricDepth(conf app.ServicePollerConfig) int {
	if conf.Metrics.Depth > 0 {
		return conf.Metrics.Depth
	}

	if conf.Metrics.Depth == 0 {
		return serviceMetricDepth
	}

	return 1
}
