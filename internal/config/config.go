package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	app "github.com/artarts36/service-navigator/internal/application"
	"github.com/artarts36/service-navigator/internal/presentation/config"
)

const serviceMetricDepth = 50
const servicePollInterval = 2 * time.Second
const imagePollInterval = 1 * time.Minute
const volumePollInterval = 1 * time.Minute

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
		Volumes struct {
			Poll app.VolumePollerConfig `yaml:"poll"`
		} `yaml:"volumes"`
	} `yaml:"backend"`
	Parameters struct {
		LogLevel string `yaml:"log_level"`
	} `yaml:"parameters"`
	Credentials struct {
		GithubToken string `yaml:"github_token"`
	} `yaml:"credentials"`
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
	if conf.Backend.Volumes.Poll.Interval == 0 || conf.Backend.Volumes.Poll.Interval < 0 {
		conf.Backend.Volumes.Poll.Interval = volumePollInterval
	}

	if conf.Frontend.AppName == "" {
		conf.Frontend.AppName = "ServiceNavigator"
	}

	if conf.Credentials.GithubToken != "" {
		conf.Credentials.GithubToken, err = resolveEnvString(conf.Credentials.GithubToken)
		if err != nil {
			panic(fmt.Sprintf("failed to resolve github token: %s", err))
		}
	}

	conf.Backend.NetworkName, err = resolveEnvString(conf.Backend.NetworkName)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve network name: %s", err))
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

func resolveEnvString(v string) (string, error) {
	if !strings.HasPrefix(v, "$") {
		return v, nil
	}

	v = strings.TrimLeft(v, "${")
	v = strings.TrimRight(v, "}")

	env, ok := os.LookupEnv(v)
	if !ok {
		return "", fmt.Errorf("environment variable %s not set", v)
	}

	return env, nil
}
