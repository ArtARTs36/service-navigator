package monitor

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type Filler interface {
	Fill(image *domain.Image, meta *datastruct.ImageMeta)
}
