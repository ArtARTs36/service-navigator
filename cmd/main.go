package main

import (
	"fmt"
	"github.com/artarts36/service-navigator/internal/http/handlers"
	"github.com/artarts36/service-navigator/internal/services"
	weburl2 "github.com/artarts36/service-navigator/internal/services/weburl"
	"github.com/docker/docker/client"
	"log"
	"net/http"
	"time"
)

type container struct {
	dockerClient *client.Client
	services     struct {
		monitor *services.Monitor
	}
	http struct {
		handlers struct {
			mainPageHandler *handlers.MainPageHandler
		}
	}
}

func main() {
	cont := initContainer()

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

func initContainer() *container {
	cont := &container{}

	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(fmt.Sprintf("Failed to create docker client: %s", err))
	}

	cont.services.monitor = services.NewMonitor(docker, weburl2.NewCompositeResolver([]weburl2.UrlResolver{
		&weburl2.NginxProxyUrlResolver{},
	}))

	cont.dockerClient = docker
	cont.http.handlers.mainPageHandler = handlers.NewMainPageHandler(cont.services.monitor)

	return cont
}
