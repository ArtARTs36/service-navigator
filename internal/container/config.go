package container

import (
	"fmt"
	"log"
	"os"

	"github.com/artarts36/service-navigator/internal/presentation/config"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Frontend config.Frontend `yaml:"frontend"`
	Backend  struct {
		NetworkName string `yaml:"network_name"`
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

	return &conf
}
