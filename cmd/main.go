package main

import (
	"fmt"
	"github.com/artarts36/service-navigator/internal/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation"
	weburl2 "github.com/artarts36/service-navigator/internal/service/filler"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/docker/docker/client"
	"github.com/tyler-sommer/stick"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

type Environment struct {
	Frontend struct {
		AppName string `yaml:"app_name"`
		Navbar  struct {
			Links []struct {
				Title string `yaml:"title"`
				Icon  string `yaml:"icon"`
				Url   string `yaml:"url"`
			} `yaml:"links"`
		} `yaml:"navbar"`
	}
	Backend struct {
		NetworkName string `yaml:"network_name"`
	}
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

	cont := initContainer(env)

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

func initContainer(env *Environment) *container {
	cont := &container{}

	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(fmt.Sprintf("Failed to create docker client: %s", err))
	}

	cont.services.monitor = monitor.NewMonitor(docker, weburl2.NewCompositeFiller([]monitor.Filler{
		&weburl2.NginxProxyUrlFiller{},
		&weburl2.VCSFiller{},
		&weburl2.DCNameFiller{},
	}))

	cont.dockerClient = docker
	cont.presentation.renderer = initRenderer(env)
	cont.http.handlers.mainPageHandler = handlers.NewMainPageHandler(cont.services.monitor, cont.presentation.renderer)

	return cont
}

func initEnvironment() *Environment {
	env := Environment{}

	log.Printf("Loading Environment from /app/service_navigator.yaml")

	yamlContent, err := os.ReadFile("/app/service_navigator.yaml")

	if err != nil {
		panic(fmt.Sprintf("Failed to read \"/app/service_navigator.yaml\": %s", err))
	}

	err = yaml.Unmarshal(yamlContent, &env)

	if err != nil {
		panic(fmt.Sprintf("failed to load Environment: %s", err))
	}

	log.Printf("Environment loaded: %v", env)

	return &env
}

func initRenderer(env *Environment) *presentation.Renderer {
	vars := map[string]stick.Value{}
	vars["_navBar"] = env.Frontend.Navbar
	vars["_appName"] = env.Frontend.AppName

	return presentation.NewRenderer("/app/templates", vars)
}
