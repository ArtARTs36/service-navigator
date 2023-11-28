package monitor

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/shared"
)

type Monitor struct {
	docker *client.Client
	filler Filler
}

type RemovedImage struct {
	Deleted  string
	Untagged string
}

type RemoveError struct {
	Error        error
	MustBeForced bool
}

func NewMonitor(docker *client.Client, filler Filler) *Monitor {
	return &Monitor{
		docker: docker,
		filler: filler,
	}
}

func (m *Monitor) Show(ctx context.Context) ([]*domain.Image, error) {
	summary, err := m.docker.ImageList(ctx, types.ImageListOptions{})

	log.Debugf("[Image][Monitor] Fetched %d images", len(summary))

	if err != nil {
		return nil, err
	}

	images := make([]*domain.Image, 0, len(summary))

	for _, image := range summary {
		if len(image.RepoTags) == 0 {
			continue
		}

		img := &domain.Image{
			ID:       image.ID,
			Size:     image.Size,
			SizeText: shared.BytesToReadableText(image.Size),
		}

		images = append(images, img)

		m.filler.Fill(img, &datastruct.ImageMeta{
			Labels:   image.Labels,
			RepoTags: image.RepoTags,
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
