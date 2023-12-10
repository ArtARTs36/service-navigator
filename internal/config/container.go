package config

import (
	"fmt"
	vlmmonitor "github.com/artarts36/service-navigator/internal/infrastructure/volume/monitor"
	"net/http"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/application"
	imgfiller "github.com/artarts36/service-navigator/internal/infrastructure/image/filler"
	imgmonitor "github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/parser"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
	"github.com/artarts36/service-navigator/internal/presentation/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation/http/middlewares"
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
	Volumes struct {
		Monitor             *vlmmonitor.Monitor
		Repository          *repository.VolumeRepository
		Poller              *application.VolumePoller
		PollRequestsChannel chan bool
	}
	Presentation struct {
		View struct {
			Renderer *view.Renderer
		}
		HTTP struct {
			Handlers struct {
				HomePageHandler      http.Handler
				ContainerKillHandler http.Handler
				ImageListHandler     http.Handler
				ImageRemoveHandler   http.Handler
				ImageRefreshHandler  http.Handler
				VolumeListHandler    http.Handler
				VolumeRefreshHandler http.Handler
			}
		}
	}
}

func InitContainer() *Container {
	return initContainerWithConfig(InitEnvironment(), InitConfig())
}

func initContainerWithConfig(env *Environment, cfg *Config) *Container {
	setupLogger(cfg)

	cont := &Container{}

	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(fmt.Sprintf("Failed to create docker client: %s", err))
	}

	imgparser := &parser.ImageParser{}

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
		filler.NewImageFiller(imgparser),
	}), cfg.Backend.NetworkName, env.CurrentContainerID)

	cont.Services.Repository = &repository.ServiceRepository{}
	cont.Services.Poller = application.NewServicePoller(
		cont.Services.Monitor,
		cont.Services.Repository,
		&cfg.Backend.Services.Poll,
	)

	cont.Images.Monitor = imgmonitor.NewMonitor(docker, imgfiller.NewCompositeFiller([]imgmonitor.Filler{
		&imgfiller.UnknownFiller{},
		&imgfiller.NameFiller{},
		imgfiller.WhenKnownImage(imgfiller.NewShortFiller(imgparser)),
		imgfiller.WhenKnownImage(&imgfiller.VCSFiller{}),
	}))
	cont.Images.Repository = &repository.ImageRepository{}
	cont.Images.Poller = application.NewImagePoller(cont.Images.Monitor, cont.Images.Repository, &cfg.Backend.Images.Poll)
	cont.Images.PollRequestsChannel = make(chan bool)

	cont.Volumes.Monitor = vlmmonitor.NewMonitor(docker)
	cont.Volumes.Repository = &repository.VolumeRepository{}
	cont.Volumes.Poller = application.NewVolumePoller(cont.Volumes.Monitor, cont.Volumes.Repository, &cfg.Backend.Volumes.Poll)
	cont.Volumes.PollRequestsChannel = make(chan bool)

	cont.DockerClient = docker
	cont.Presentation.View.Renderer = initRenderer(env, cfg)

	cont.Presentation.HTTP.Handlers.HomePageHandler = middlewares.NewLogMiddleware(
		handlers.NewHomePageHandler(cont.Services.Repository, cont.Presentation.View.Renderer),
	)
	cont.Presentation.HTTP.Handlers.ContainerKillHandler = middlewares.NewLogMiddleware(
		handlers.NewContainerKillHandler(
			cont.Services.Monitor,
			cont.Presentation.View.Renderer,
		),
	)
	cont.Presentation.HTTP.Handlers.ImageListHandler = middlewares.NewLogMiddleware(
		handlers.NewImageListHandler(
			cont.Images.Repository,
			cont.Presentation.View.Renderer,
		),
	)
	cont.Presentation.HTTP.Handlers.ImageRemoveHandler = middlewares.NewLogMiddleware(
		handlers.NewImageRemoveHandler(
			cont.Images.Monitor,
			cont.Presentation.View.Renderer,
		),
	)
	cont.Presentation.HTTP.Handlers.ImageRefreshHandler = middlewares.NewLogMiddleware(
		handlers.NewImageRefreshHandler(cont.Images.PollRequestsChannel),
	)
	cont.Presentation.HTTP.Handlers.VolumeListHandler = middlewares.NewLogMiddleware(
		handlers.NewVolumeListHandler(cont.Volumes.Repository, cont.Presentation.View.Renderer),
	)
	cont.Presentation.HTTP.Handlers.VolumeRefreshHandler = middlewares.NewLogMiddleware(
		handlers.NewVolumeRefreshHandler(cont.Volumes.PollRequestsChannel),
	)

	return cont
}

func setupLogger(cfg *Config) {
	level, err := log.ParseLevel(cfg.Parameters.LogLevel)
	if err != nil {
		log.Warnf("log level \"%s\" unsupported", cfg.Parameters.LogLevel)

		level = log.DebugLevel
	}

	log.SetLevel(level)

	log.Debugf("setup log level \"%s\"", level)
}
