package application

import (
	"context"
	"log"
	"time"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
)

type Metrics struct {
	Depth      int  `yaml:"depth"`
	OnlyUnique bool `yaml:"only_unique"`
}

type Poller struct {
	monitor  *monitor.Monitor
	services *repository.ServiceRepository
	interval time.Duration
	metrics  *Metrics
}

func NewPoller(
	monitor *monitor.Monitor,
	serviceRepo *repository.ServiceRepository,
	interval time.Duration,
	metrics *Metrics,
) *Poller {
	return &Poller{
		monitor:  monitor,
		services: serviceRepo,
		interval: interval,
		metrics:  metrics,
	}
}

func (p *Poller) Poll() {
	for {
		statuses, err := p.monitor.Show(context.Background())

		if err != nil {
			log.Printf("[Poller] Failed to load statuses: %s", err)

			continue
		}

		existsServicesMap := make(map[string]*domain.Service)

		for _, service := range p.services.All() {
			existsServicesMap[service.ContainerID] = service
		}

		newServicesList := make([]*domain.Service, 0, len(statuses))

		for _, status := range statuses {
			service, serviceExists := existsServicesMap[status.ContainerID]

			if !serviceExists {
				service = domain.NewService(p.metrics.Depth, p.metrics.OnlyUnique)
			}

			service.Name = status.Name
			service.WebURL = status.WebURL
			service.Status = status.Status
			service.VCS = status.VCS
			service.ContainerID = status.ContainerID
			service.MemoryHistory.Push(status.Memory)
			service.Self = status.Self

			newServicesList = append(newServicesList, service)
		}

		p.services.Set(newServicesList)

		log.Printf("[Poller] loaded %d statuses", len(statuses))

		time.Sleep(p.interval)
	}
}
