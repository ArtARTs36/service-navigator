package application

import (
	"context"
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

func (p *VolumePoller) Poll(ctx context.Context, reqs chan bool) {
	tick := time.NewTicker(p.config.Interval).C
	ticked := false

	for {
		select {
		case <-ctx.Done():
			log.Debug("[Volume][Poller] Stopped")
			return
		case <-reqs:
			log.Debug("[Volume][Poller] Given user request")

			err := p.poll()
			if err != nil {
				log.Warnf("[Volume][Poller] Failed to load statuses: %s", err)
				continue
			}
		case <-tick:
			err := p.poll()
			if err != nil {
				log.Warnf("[Volume][Poller] Failed to load statuses: %s", err)
				continue
			}

			log.Infof("[Volume][Poller] sleep %.2f seconds", p.config.Interval.Seconds())
		default:
			if !ticked {
				err := p.poll()
				if err != nil {
					log.Debugf("[Volume][Poller] Failed to load statuses: %s", err)
					continue
				}

				log.Debugf("[Volume][Poller] sleep %.2f seconds", p.config.Interval.Seconds())
				ticked = true
			}
		}
	}
}

func (p *VolumePoller) poll() error {
	images, err := p.monitor.Show(context.Background())

	if err != nil {
		log.Warnf("[Volume][Poller] Failed to load volumes: %s", err)

		return err
	}

	log.Debugf("[Volume][Poller] loaded %d volumes", len(images))

	p.volumes.Set(images)

	return nil
}
