package monitor

import (
	"context"
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

type Monitor struct {
	docker *client.Client
}

func NewMonitor(docker *client.Client) *Monitor {
	return &Monitor{
		docker: docker,
	}
}

func (m *Monitor) Show(ctx context.Context) ([]*domain.Volume, error) {
	list, err := m.docker.VolumeList(ctx, volume.ListOptions{})

	log.Debugf("[Volume][Monitor] Fetched %d volumes", len(list.Volumes))

	if err != nil {
		return nil, err
	}

	volumes := make([]*domain.Volume, 0, len(list.Volumes))

	for _, vol := range list.Volumes {
		var size *int64

		if vol.UsageData != nil {
			size = &vol.UsageData.Size
		}

		volumes = append(volumes, &domain.Volume{
			Name:      vol.Name,
			Driver:    vol.Driver,
			Size:      size,
			CreatedAt: vol.CreatedAt,
		})
	}

	return volumes, nil
}
