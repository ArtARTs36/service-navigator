package config

import (
	"fmt"
	"time"

	poller "github.com/artarts36/service-navigator/internal/application"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
	handlers2 "github.com/artarts36/service-navigator/internal/presentation/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation/view"
	"github.com/docker/docker/client"
)

const servicePollInterval = 1 * time.Second

type Container struct {
	DockerClient *client.Client
	Services     struct {
		Monitor    *monitor.Monitor
		Repository *repository.ServiceRepository
		Poller     *poller.Poller
	}
	HTTP struct {
		Handlers struct {
			HomePageHandler      *handlers2.HomePageHandler
			ContainerKIllHandler *handlers2.ContainerKillHandler
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
	}), conf.Backend.NetworkName, env.CurrentContainerID)

	cont.Services.Repository = &repository.ServiceRepository{}
	cont.Services.Poller = poller.NewPoller(
		cont.Services.Monitor,
		cont.Services.Repository,
		servicePollInterval,
		&conf.Backend.Metrics,
	)

	cont.DockerClient = docker
	cont.Presentation.Renderer = initRenderer(env, conf)
	cont.HTTP.Handlers.HomePageHandler = handlers2.NewHomePageHandler(cont.Services.Repository, cont.Presentation.Renderer)
	cont.HTTP.Handlers.ContainerKIllHandler = handlers2.NewContainerKillHandler(
		cont.Services.Monitor,
		cont.Presentation.Renderer,
	)

	return cont
}
