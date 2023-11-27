package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type NameFiller struct {
}

func (f *NameFiller) Fill(image *domain.Image, meta *datastruct.ImageMeta) {
	if image.Unknown {
		image.Name = image.ID

		return
	}

	if len(meta.RepoTags) > 0 {
		image.Name = meta.RepoTags[0]
	}
}
