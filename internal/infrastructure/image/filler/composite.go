package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type CompositeFiller struct {
	fillers []monitor.Filler
}

func NewCompositeFiller(fillers []monitor.Filler) monitor.Filler {
	return &CompositeFiller{
		fillers: fillers,
	}
}

func (r *CompositeFiller) Fill(image *domain.Image, meta *datastruct.ImageMeta) {
	for _, filler := range r.fillers {
		filler.Fill(image, meta)
	}
}
