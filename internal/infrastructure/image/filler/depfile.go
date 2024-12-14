package filler

import (
	"context"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/dep"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	log "github.com/sirupsen/logrus"
)

type DepFileFiller struct {
	depFetcher dep.Fetcher
}

func NewDepFileFiller(fetcher dep.Fetcher) *DepFileFiller {
	return &DepFileFiller{
		depFetcher: fetcher,
	}
}

func (f *DepFileFiller) Fill(image *domain.Image, _ *datastruct.ImageMeta) {
	if image.VCS == nil {
		return
	}

	depFile, err := f.depFetcher.Fetch(context.Background(), image)
	if err != nil {
		log.Warn("failed to fetch dependency file from repo", err, map[string]interface{}{
			"image_name": image.Name,
			"vcs_url":    image.VCS.URL,
		})
		return
	}

	image.DepFiles = depFile
}
