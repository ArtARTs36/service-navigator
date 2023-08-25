package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/artarts36/service-navigator/internal/container"
	"github.com/artarts36/service-navigator/internal/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation"
	weburl2 "github.com/artarts36/service-navigator/internal/service/filler"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/docker/docker/client"
	"github.com/tyler-sommer/stick"
)

const httpReadTimeout = 3 * time.Second

func main() {
	env := container.InitEnvironment()
	conf := container.InitConfig()

	cont := initContainer(env, conf)

	mux := http.NewServeMux()
	mux.Handle("/", cont.HTTP.Handlers.HomePageHandler)
	mux.Handle("/containers/kill", cont.HTTP.Handlers.ContainerKIllHandler)

	fs := http.FileServer(http.Dir("/app/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	hServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: httpReadTimeout,
	}

	log.Print("Listening...")

	err := hServer.ListenAndServe()
	if err != nil {
		log.Printf("Failed listeing: %s", err)

		return
	}
}

func initContainer(env *container.Environment, conf *container.Config) *container.Container {
	cont := &container.Container{}

	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(fmt.Sprintf("Failed to create docker client: %s", err))
	}

	cont.Services.Monitor = monitor.NewMonitor(docker, weburl2.NewCompositeFiller([]monitor.Filler{
		&weburl2.NginxProxyURLFiller{},
		&weburl2.VCSFiller{},
		&weburl2.DCNameFiller{},
	}), conf.Backend.NetworkName, env.CurrentContainerID)

	cont.DockerClient = docker
	cont.Presentation.Renderer = initRenderer(env, conf)
	cont.HTTP.Handlers.HomePageHandler = handlers.NewHomePageHandler(cont.Services.Monitor, cont.Presentation.Renderer)
	cont.HTTP.Handlers.ContainerKIllHandler = handlers.NewContainerKillHandler(
		cont.Services.Monitor,
		cont.Presentation.Renderer,
	)

	return cont
}

func initRenderer(env *container.Environment, conf *container.Config) *presentation.Renderer {
	vars := map[string]stick.Value{}
	vars["_navBar"] = conf.Frontend.Navbar
	vars["_appName"] = conf.Frontend.AppName
	vars["_username"] = env.User

	return presentation.NewRenderer("/app/templates", vars)
}
