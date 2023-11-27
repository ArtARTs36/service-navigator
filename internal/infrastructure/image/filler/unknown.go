package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type UnknownFiller struct {
}

func (f *UnknownFiller) Fill(image *domain.Image, meta *datastruct.ImageMeta) {
	image.Unknown = len(meta.RepoTags) > 0 && meta.RepoTags[0] == "<none>:<none>"
}
