package application

import (
	"context"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
	"log"
	"time"
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

func (p *ImagePoller) Poll() {
	for {
		images, err := p.monitor.Show(context.Background())

		if err != nil {
			log.Printf("[Image][Poll] Failed to load statuses: %s", err)

			continue
		}

		log.Printf("[Image][Poll] loaded %d images", len(images))
		log.Printf("[Image][Poll] sleep %.2f seconds", p.config.Interval.Seconds())

		p.images.Set(images)

		time.Sleep(p.config.Interval)
	}
}
