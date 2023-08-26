package config

import (
	"fmt"

	"github.com/artarts36/service-navigator/internal/application"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
	"github.com/artarts36/service-navigator/internal/presentation/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation/view"
	"github.com/docker/docker/client"
)

type Container struct {
	DockerClient *client.Client
	Services     struct {
		Monitor    *monitor.Monitor
		Repository *repository.ServiceRepository
		Poller     *application.Poller
	}
	HTTP struct {
		Handlers struct {
			HomePageHandler      *handlers.HomePageHandler
			ContainerKIllHandler *handlers.ContainerKillHandler
		}
	}
	Presentation struct {
		Renderer *view.Renderer
	}
}

func InitContainer() *Container {
	return initContainerWithConfig(InitEnvironment(), InitConfig())
}

func initContainerWithConfig(env *Environment, conf *Config) *Container {
	cont := &Container{}

	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(fmt.Sprintf("Failed to create docker client: %s", err))
	}

	cont.Services.Monitor = monitor.NewMonitor(docker, filler.NewCompositeFiller([]monitor.Filler{
		&filler.NginxProxyURLFiller{},
		&filler.VCSFiller{},
		&filler.DCNameFiller{},
		&filler.MemoryFiller{},
		&filler.CPUFiller{},
	}), conf.Backend.NetworkName, env.CurrentContainerID)

	cont.Services.Repository = &repository.ServiceRepository{}
	cont.Services.Poller = application.NewPoller(
		cont.Services.Monitor,
		cont.Services.Repository,
		&conf.Backend.Poll,
	)

	cont.DockerClient = docker
	cont.Presentation.Renderer = initRenderer(env, conf)
	cont.HTTP.Handlers.HomePageHandler = handlers.NewHomePageHandler(cont.Services.Repository, cont.Presentation.Renderer)
	cont.HTTP.Handlers.ContainerKIllHandler = handlers.NewContainerKillHandler(
		cont.Services.Monitor,
		cont.Presentation.Renderer,
	)

	return cont
}
