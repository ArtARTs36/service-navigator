package main

import (
	"fmt"
	"github.com/artarts36/service-navigator/internal/http/handlers"
	weburl2 "github.com/artarts36/service-navigator/internal/service/filler"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/docker/docker/client"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"time"
)

type environment struct {
	AppName     string `envconfig:"APP_NAME" default:"ServiceNavigator"`
	NetworkName string `envconfig:"NETWORK_NAME" required:"true"`
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
}

func main() {
	var env environment

	err := envconfig.Process("SERVICE_NAVIGATOR", &env)

	if err != nil {
		panic(fmt.Sprintf("failed to load environment: %s", err))
	}

	cont := initContainer(&env)

	mux := http.NewServeMux()
	mux.Handle("/", cont.http.handlers.mainPageHandler)
	hServer := &http.Server{
		Addr:        ":9100",
		Handler:     mux,
		ReadTimeout: 3 * time.Second,
	}

	log.Print("Listening...")

	err = hServer.ListenAndServe()
	if err != nil {
		return
	}
}

func initContainer(env *environment) *container {
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
	cont.http.handlers.mainPageHandler = handlers.NewMainPageHandler(cont.services.monitor, env.AppName, env.NetworkName)

	return cont
}
