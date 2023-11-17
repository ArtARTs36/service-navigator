package monitor

import (
	"context"
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/shared"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"strings"
)

type Monitor struct {
	docker *client.Client
}

type RemovedImage struct {
	Deleted  string
	Untagged string
}

type RemoveError struct {
	Error        error
	MustBeForced bool
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

func (m *Monitor) Remove(ctx context.Context, imageID string, force bool) ([]*RemovedImage, *RemoveError) {
	items, err := m.docker.ImageRemove(ctx, imageID, types.ImageRemoveOptions{
		Force: force,
	})
	if err != nil {
		return nil, &RemoveError{
			Error:        err,
			MustBeForced: strings.Contains(err.Error(), "must be forced"),
		}
	}

	removed := make([]*RemovedImage, 0, len(items))

	for _, item := range items {
		removed = append(removed, &RemovedImage{
			Deleted:  item.Deleted,
			Untagged: item.Untagged,
		})
	}

	return removed, nil
}
