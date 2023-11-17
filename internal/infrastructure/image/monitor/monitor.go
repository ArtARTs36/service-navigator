package monitor

import (
	"context"
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/shared"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

type Monitor struct {
	docker *client.Client
}

func NewMonitor(docker *client.Client) *Monitor {
	return &Monitor{
		docker: docker,
	}
}

func (m *Monitor) Show(ctx context.Context) ([]*domain.Image, error) {
	summary, err := m.docker.ImageList(ctx, types.ImageListOptions{})

	log.Printf("[Image][Monitor] Fetched %d images", len(summary))

	if err != nil {
		return nil, err
	}

	images := make([]*domain.Image, 0, len(summary))

	for _, image := range summary {
		if len(image.RepoTags) == 0 {
			continue
		}

		name := image.RepoTags[0]
		if name == "<none>:<none>" {
			continue
		}

		images = append(images, &domain.Image{
			ID:       image.ID,
			Name:     name,
			Size:     image.Size,
			SizeText: shared.BytesToReadableText(image.Size),
		})
	}

	return images, nil
}
