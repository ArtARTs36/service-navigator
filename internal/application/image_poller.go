package application

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/repository"
)

type ImagePollerConfig struct {
	Interval time.Duration `yaml:"interval"`

	ScanRepo ScanRepoConfig `yaml:"scan_repo"`
}

type ScanRepoConfig struct {
	Lang bool `yaml:"lang"`
	Deps bool `yaml:"deps"`
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

func (p *ImagePoller) Poll(ctx context.Context, reqs chan bool) {
	tick := time.NewTicker(p.config.Interval).C

	p.poll()

	for {
		select {
		case <-ctx.Done():
			log.Print("[Image][Poller] Stopped")
			return
		case <-reqs:
			log.Print("[Image][Poller] Given user request")

			p.poll()
		case <-tick:
			p.poll()

			log.Printf("[Image][Poller] sleep %.2f seconds", p.config.Interval.Seconds())
		}
	}
}

func (p *ImagePoller) poll() {
	images, err := p.monitor.Show(context.Background())
	if err != nil {
		log.Printf("[Image][Poller] Failed to load images: %s", err)
		return
	}

	log.Printf("[Image][Poller] loaded %d volumes", len(images))

	p.images.Set(images)
}

func (s *ScanRepoConfig) Enabled() bool {
	return s.Lang || s.Deps
}
