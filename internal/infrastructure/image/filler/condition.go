package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/monitor"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type Condition func(image *domain.Image, meta *datastruct.ImageMeta) bool

type ConditionFiller struct {
	filler    monitor.Filler
	condition Condition
}

func NewConditionFiller(filler monitor.Filler, condition Condition) monitor.Filler {
	return &ConditionFiller{
		filler:    filler,
		condition: condition,
	}
}

func WhenKnownImage(filler monitor.Filler) monitor.Filler {
	return NewConditionFiller(filler, func(image *domain.Image, _ *datastruct.ImageMeta) bool {
		return !image.Unknown
	})
}

func (f *ConditionFiller) Fill(image *domain.Image, meta *datastruct.ImageMeta) {
	if f.condition(image, meta) {
		f.filler.Fill(image, meta)
	}
}
