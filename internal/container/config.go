package container

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/artarts36/service-navigator/internal/infrastructure/search"
)

type Config struct {
	Frontend struct {
		AppName string `yaml:"app_name"`
		Navbar  struct {
			Links []struct {
				Title string `yaml:"title"`
				Icon  string `yaml:"icon"`
				URL   string `yaml:"url"`
			} `yaml:"links"`
			Profile struct {
				Links []struct {
					Title string `yaml:"title"`
					Icon  string `yaml:"icon"`
					URL   string `yaml:"url"`
				} `yaml:"links"`
			} `yaml:"profile"`
			Search struct {
				Providers []search.Provider `yaml:"providers"`
			} `yaml:"search"`
		} `yaml:"navbar"`
	} `yaml:"frontend"`
	Backend struct {
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

	conf.Frontend.Navbar.Search.Providers = search.ResolveProviders(conf.Frontend.Navbar.Search.Providers)

	log.Printf("Config loaded: %v", conf)

	return &conf
}
