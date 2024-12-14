package config

import (
	"fmt"
	"net/http"

	"github.com/artarts36/depexplorer/pkg/github"
	depRepo "github.com/artarts36/depexplorer/pkg/repository"

	"github.com/artarts36/service-navigator/internal/infrastructure/image/dep"
	githubClient "github.com/google/go-github/v67/github"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/application"
	imgfiller "github.com/artarts36/service-navigator/internal/infrastructure/image/filler"
	imgmonitor "github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/parser"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
	vlmmonitor "github.com/artarts36/service-navigator/internal/infrastructure/volume/monitor"
	"github.com/artarts36/service-navigator/internal/presentation/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation/http/middlewares"
	"github.com/artarts36/service-navigator/internal/presentation/view"
)

type Container struct {
	Services struct {
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
				ImageShowHandler     http.Handler
				ImageRemoveHandler   http.Handler
				ImageRefreshHandler  http.Handler
				VolumeListHandler    http.Handler
				VolumeRefreshHandler http.Handler
			}
		}
	}

	Clients struct {
		Github *githubClient.Client
		Docker *client.Client
	}
}

func InitContainer() (*Container, error) {
	return initContainerWithConfig(InitEnvironment(), InitConfig())
}

func initContainerWithConfig(env *Environment, cfg *Config) (*Container, error) { //nolint:funlen // not need
	setupLogger(cfg)

	cont := &Container{}

	if err := cont.initClients(cfg); err != nil {
		return nil, fmt.Errorf("failed to init clients: %w", err)
	}

	imgparser := &parser.ImageParser{}

	cont.Services.Monitor = monitor.NewMonitor(cont.Clients.Docker, filler.NewCompositeFiller([]monitor.Filler{
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

	depFileFetcher := cont.resolveDepFetcher()

	imgFillers := []imgmonitor.Filler{
		&imgfiller.UnknownFiller{},
		&imgfiller.NameFiller{},
		imgfiller.WhenKnownImage(imgfiller.NewShortFiller(imgparser)),
		imgfiller.WhenKnownImage(&imgfiller.VCSFiller{}),
		&imgfiller.CreatedFiller{},
	}
	if cfg.Backend.Images.Poll.ScanRepo.Enabled() {
		imgFillers = append(imgFillers, imgfiller.NewDepFileFiller(depFileFetcher))

		if cfg.Backend.Images.Poll.ScanRepo.Lang {
			imgFillers = append(imgFillers, imgfiller.NewLanguageFiller())
		}
	}

	cont.Images.Monitor = imgmonitor.NewMonitor(cont.Clients.Docker, imgfiller.NewCompositeFiller(imgFillers))
	cont.Images.Repository = &repository.ImageRepository{}
	cont.Images.Poller = application.NewImagePoller(cont.Images.Monitor, cont.Images.Repository, &cfg.Backend.Images.Poll)
	cont.Images.PollRequestsChannel = make(chan bool)

	cont.Volumes.Monitor = vlmmonitor.NewMonitor(cont.Clients.Docker)
	cont.Volumes.Repository = &repository.VolumeRepository{}
	cont.Volumes.Poller = application.NewVolumePoller(
		cont.Volumes.Monitor,
		cont.Volumes.Repository,
		&cfg.Backend.Volumes.Poll,
	)
	cont.Volumes.PollRequestsChannel = make(chan bool)

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
			cfg.Frontend.Pages.Images,
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
	cont.Presentation.HTTP.Handlers.ImageShowHandler = middlewares.NewLogMiddleware(
		handlers.NewImageShowHandler(cont.Presentation.View.Renderer, cont.Images.Repository),
	)

	return cont, nil
}

func (c *Container) resolveDepFetcher() dep.Fetcher {
	logger := func(s string, m map[string]interface{}) {
		log.Info(s, m)
	}

	return dep.NewStoreableFetcher(dep.NewClientFetcher(map[string]depRepo.Explorer{
		"github": github.NewExplorer(c.Clients.Github, logger),
	}), dep.NewMemoryFileStore())
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
