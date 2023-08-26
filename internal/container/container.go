package container

import (
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	handlers2 "github.com/artarts36/service-navigator/internal/presentation/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation/view"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/docker/docker/client"
)

type Container struct {
	DockerClient *client.Client
	Services     struct {
		Monitor    *monitor.Monitor
		Repository *repository.ServiceRepository
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
