package container

import (
	"github.com/artarts36/service-navigator/internal/http/handlers"
	"github.com/artarts36/service-navigator/internal/presentation"
	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/artarts36/service-navigator/internal/service/repository"
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
			HomePageHandler      *handlers.HomePageHandler
			ContainerKIllHandler *handlers.ContainerKillHandler
		}
	}
	Presentation struct {
		Renderer *presentation.Renderer
	}
}
