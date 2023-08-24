package main

import (
	"fmt"
	"github.com/artarts36/service-navigator/internal/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation"
	"github.com/artarts36/service-navigator/internal/search"
	weburl2 "github.com/artarts36/service-navigator/internal/service/filler"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/docker/docker/client"
	"github.com/kelseyhightower/envconfig"
	"github.com/tyler-sommer/stick"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Frontend struct {
		AppName string `yaml:"app_name"`
		Navbar  struct {
			Links []struct {
				Title string `yaml:"title"`
				Icon  string `yaml:"icon"`
				Url   string `yaml:"url"`
			} `yaml:"links"`
			Profile struct {
				Links []struct {
					Title string `yaml:"title"`
					Icon  string `yaml:"icon"`
					Url   string `yaml:"url"`
				} `yaml:"links"`
			} `yaml:"profile"`
			Search struct {
				Providers     []search.Provider `yaml:"providers"`
				FirstProvider search.Provider
			} `yaml:"search"`
		} `yaml:"navbar"`
	}
	Backend struct {
		NetworkName string `yaml:"network_name"`
	}
}

type Environment struct {
	User string `default:"developer"`
}

type container struct {
	dockerClient *client.Client
	services     struct {
		monitor *monitor.Monitor
	}
	http struct {
		handlers struct {
			mainPageHandler *handlers.MainPageHandler
		}
	}
	presentation struct {
		renderer *presentation.Renderer
	}
}

func main() {
	env := initEnvironment()
	conf := initConfig()

	cont := initContainer(env, conf)

	mux := http.NewServeMux()
	mux.Handle("/", cont.http.handlers.mainPageHandler)
	hServer := &http.Server{
		Addr:        ":9100",
		Handler:     mux,
		ReadTimeout: 3 * time.Second,
	}

	log.Print("Listening...")

	err := hServer.ListenAndServe()
	if err != nil {
		return
	}
}

func initContainer(env *Environment, conf *Config) *container {
	cont := &container{}

	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(fmt.Sprintf("Failed to create docker client: %s", err))
	}

	cont.services.monitor = monitor.NewMonitor(docker, weburl2.NewCompositeFiller([]monitor.Filler{
		&weburl2.NginxProxyUrlFiller{},
		&weburl2.VCSFiller{},
		&weburl2.DCNameFiller{},
	}), conf.Backend.NetworkName)

	cont.dockerClient = docker
	cont.presentation.renderer = initRenderer(env, conf)
	cont.http.handlers.mainPageHandler = handlers.NewMainPageHandler(cont.services.monitor, cont.presentation.renderer)

	return cont
}

func initConfig() *Config {
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

func initRenderer(env *Environment, conf *Config) *presentation.Renderer {
	vars := map[string]stick.Value{}
	vars["_navBar"] = conf.Frontend.Navbar
	vars["_appName"] = conf.Frontend.AppName
	vars["_username"] = env.User

	return presentation.NewRenderer("/app/templates", vars)
}

func initEnvironment() *Environment {
	var env Environment

	log.Print("Loading environment")

	err := envconfig.Process("", &env)

	if err != nil {
		panic(fmt.Sprintf("Failed to load environment: %s", err))
	}

	log.Printf("Environment loaded: %v", env)

	return &env
}
