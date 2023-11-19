package application

import (
	"context"
	"log"
	"time"

	"github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
)

type ImagePollerConfig struct {
	Interval time.Duration `yaml:"interval"`
}

type ImagePoller struct {
	monitor *monitor.Monitor
	images  *repository.ImageRepository
	config  *ImagePollerConfig
}

func NewImagePoller(
	monitor *monitor.Monitor,
	imgrepo *repository.ImageRepository,
	config *ImagePollerConfig,
) *ImagePoller {
	return &ImagePoller{
		monitor: monitor,
		images:  imgrepo,
		config:  config,
	}
}

func (p *ImagePoller) Poll(reqs chan bool) {
	tick := time.Tick(p.config.Interval)
	ticked := false

	for {
		select {
		case <-reqs:
			log.Print("[Image][Poll] Given user request")

			err := p.poll()
			if err != nil {
				log.Printf("[Image][Poll] Failed to load statuses: %s", err)
				continue
			}
		case <-tick:
			err := p.poll()
			if err != nil {
				log.Printf("[Image][Poll] Failed to load statuses: %s", err)
				continue
			}

			log.Printf("[Image][Poll] sleep %.2f seconds", p.config.Interval.Seconds())
		default:
			if !ticked {
				err := p.poll()
				if err != nil {
					log.Printf("[Image][Poll] Failed to load statuses: %s", err)
					continue
				}

				log.Printf("[Image][Poll] sleep %.2f seconds", p.config.Interval.Seconds())
				ticked = true
			}
		}
	}
}

func (p *ImagePoller) poll() error {
	images, err := p.monitor.Show(context.Background())

	if err != nil {
		log.Printf("[Image][Poll] Failed to load images: %s", err)

		return err
	}

	log.Printf("[Image][Poll] loaded %d images", len(images))

	p.images.Set(images)

	return nil
}
