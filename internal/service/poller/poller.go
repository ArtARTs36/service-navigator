package poller

import (
	"context"
	"log"
	"time"

	"github.com/artarts36/service-navigator/internal/service/monitor"
	"github.com/artarts36/service-navigator/internal/service/repository"
)

type Poller struct {
	monitor  *monitor.Monitor
	services *repository.ServiceRepository
	interval time.Duration
}

func NewPoller(monitor *monitor.Monitor, serviceRepo *repository.ServiceRepository, interval time.Duration) *Poller {
	return &Poller{
		monitor:  monitor,
		services: serviceRepo,
		interval: interval,
	}
}

func (p *Poller) Poll() {
	for {
		services, err := p.monitor.Show(context.Background())

		if err != nil {
			log.Printf("[Poller] Failed to load services: %s", err)

			continue
		}

		p.services.Set(services)

		log.Printf("[Poller] loaded %d services", len(services))

		time.Sleep(p.interval)
	}
}
