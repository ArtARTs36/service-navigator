package application

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"github.com/artarts36/service-navigator/internal/infrastructure/volume/monitor"
)

type VolumePollerConfig struct {
	Interval time.Duration `yaml:"interval"`
}

type VolumePoller struct {
	monitor *monitor.Monitor
	volumes *repository.VolumeRepository
	config  *VolumePollerConfig
}

func NewVolumePoller(
	monitor *monitor.Monitor,
	vlmrepo *repository.VolumeRepository,
	config *VolumePollerConfig,
) *VolumePoller {
	return &VolumePoller{
		monitor: monitor,
		volumes: vlmrepo,
		config:  config,
	}
}

func (p *VolumePoller) Poll(ctx context.Context, wg *sync.WaitGroup, reqs chan bool) {
	defer wg.Done()

	tick := time.NewTicker(p.config.Interval).C
	ticked := false

	for {
		select {
		case <-ctx.Done():
			log.Print("[Volume][Poller] Stopped")
			return
		case <-reqs:
			log.Print("[Volume][Poller] Given user request")

			err := p.poll()
			if err != nil {
				log.Printf("[Volume][Poller] Failed to load statuses: %s", err)
				continue
			}
		case <-tick:
			err := p.poll()
			if err != nil {
				log.Printf("[Volume][Poller] Failed to load statuses: %s", err)
				continue
			}

			log.Printf("[Volume][Poller] sleep %.2f seconds", p.config.Interval.Seconds())
		default:
			if !ticked {
				err := p.poll()
				if err != nil {
					log.Printf("[Volume][Poller] Failed to load statuses: %s", err)
					continue
				}

				log.Printf("[Volume][Poller] sleep %.2f seconds", p.config.Interval.Seconds())
				ticked = true
			}
		}
	}
}

func (p *VolumePoller) poll() error {
	images, err := p.monitor.Show(context.Background())

	if err != nil {
		log.Printf("[Volume][Poller] Failed to load volumes: %s", err)

		return err
	}

	log.Printf("[Volume][Poller] loaded %d volumes", len(images))

	p.volumes.Set(images)

	return nil
}
