package monitor

import (
	"context"

	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/shared"
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
		vlm := &domain.Volume{
			Name:      vol.Name,
			Driver:    vol.Driver,
			CreatedAt: vol.CreatedAt,
		}

		if vol.UsageData != nil {
			vlm.Size = shared.Int64ToPtr(vol.UsageData.Size)
		}

		volumes = append(volumes, vlm)
	}

	return volumes, nil
}
