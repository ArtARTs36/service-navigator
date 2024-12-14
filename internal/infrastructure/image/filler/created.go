package filler

import (
	"time"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type CreatedFiller struct{}

func (c *CreatedFiller) Fill(image *domain.Image, meta *datastruct.ImageMeta) {
	image.CreatedAt = time.Unix(meta.Created, 0)
}
