package config

import (
	"fmt"

	"github.com/docker/docker/client"

	"github.com/artarts36/service-navigator/internal/application"
	imgmonitor "github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
	"github.com/artarts36/service-navigator/internal/presentation/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation/view"
)

type Container struct {
	DockerClient *client.Client
	Services     struct {
		Monitor    *monitor.Monitor
		Repository *repository.ServiceRepository
		Poller     *application.ServicePoller
	}
	Images struct {
		Monitor             *imgmonitor.Monitor
		Repository          *repository.ImageRepository
		Poller              *application.ImagePoller
		PollRequestsChannel chan bool
	}
	HTTP struct {
		Handlers struct {
			HomePageHandler      *handlers.HomePageHandler
			ContainerKillHandler *handlers.ContainerKillHandler
			ImageListHandler     *handlers.ImageListHandler
			ImageRemoveHandler   *handlers.ImageRemoveHandler
			ImageRefreshHandler  *handlers.ImageRefreshHandler
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
		filler.NewOrFiller([]monitor.Filler{
			filler.NewNginxProxyURLFiller(),
			filler.NewPublicPortFiller(),
		}),
		&filler.VCSFiller{},
		filler.NewOrFiller([]monitor.Filler{
			&filler.DCNameFiller{},
			&filler.SwarmNameFiller{},
		}),
		&filler.MemoryFiller{},
		&filler.CPUFiller{},
		&filler.ImageFiller{},
	}), conf.Backend.NetworkName, env.CurrentContainerID)

	cont.Services.Repository = &repository.ServiceRepository{}
	cont.Services.Poller = application.NewServicePoller(
		cont.Services.Monitor,
		cont.Services.Repository,
		&conf.Backend.Services.Poll,
	)

	cont.Images.Monitor = imgmonitor.NewMonitor(docker)
	cont.Images.Repository = &repository.ImageRepository{}
	cont.Images.Poller = application.NewImagePoller(cont.Images.Monitor, cont.Images.Repository, &conf.Backend.Images.Poll)
	cont.Images.PollRequestsChannel = make(chan bool)

	cont.DockerClient = docker
	cont.Presentation.Renderer = initRenderer(env, conf)
	cont.HTTP.Handlers.HomePageHandler = handlers.NewHomePageHandler(cont.Services.Repository, cont.Presentation.Renderer)
	cont.HTTP.Handlers.ContainerKillHandler = handlers.NewContainerKillHandler(
		cont.Services.Monitor,
		cont.Presentation.Renderer,
	)
	cont.HTTP.Handlers.ImageListHandler = handlers.NewImageListHandler(
		cont.Images.Repository,
		cont.Presentation.Renderer,
	)
	cont.HTTP.Handlers.ImageRemoveHandler = handlers.NewImageRemoveHandler(
		cont.Images.Monitor,
		cont.Presentation.Renderer,
	)
	cont.HTTP.Handlers.ImageRefreshHandler = handlers.NewImageRefreshHandler(cont.Images.PollRequestsChannel)

	return cont
}
