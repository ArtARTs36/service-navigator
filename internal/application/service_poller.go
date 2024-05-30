package application

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
)

type ServicePollerConfig struct {
	Interval time.Duration `yaml:"interval"`
	Metrics  struct {
		Depth      int  `yaml:"depth"`
		OnlyUnique bool `yaml:"only_unique"`
	} `yaml:"metrics"`
}

type ServicePoller struct {
	monitor  *monitor.Monitor
	services *repository.ServiceRepository
	config   *ServicePollerConfig
}

func NewServicePoller(
	monitor *monitor.Monitor,
	serviceRepo *repository.ServiceRepository,
	config *ServicePollerConfig,
) *ServicePoller {
	return &ServicePoller{
		monitor:  monitor,
		services: serviceRepo,
		config:   config,
	}
}

func (p *ServicePoller) Poll(ctx context.Context) {
	tick := time.NewTicker(p.config.Interval).C

	for {
		select {
		case <-ctx.Done():
			log.Print("[Service][Poller] Stopped")
			return
		case <-tick:
			statuses, err := p.monitor.Show(context.Background())

			if err != nil {
				log.Printf("[Service][Poller] Failed to load statuses: %s", err)

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
					service = domain.NewService(p.config.Metrics.Depth, p.config.Metrics.OnlyUnique)
				}

				service.Name = status.Name
				service.WebURL = status.WebURL
				service.Status = status.Status
				service.VCS = status.VCS
				service.ContainerID = status.ContainerID
				service.MemoryHistory.Push(status.Memory)
				service.CPUHistory.Push(status.CPU)
				service.Self = status.Self
				service.Image = status.Image

				newServicesList = append(newServicesList, service)
			}

			p.services.Set(newServicesList)

			log.Printf("[Service][Poller] loaded %d statuses", len(statuses))
			log.Printf("[Service][Poller] sleep %.2f seconds", p.config.Interval.Seconds())

			time.Sleep(p.config.Interval)
		}
	}
}
