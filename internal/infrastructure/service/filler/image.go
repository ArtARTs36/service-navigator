package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/parser"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type ImageFiller struct {
	parser *parser.ImageParser
}

func NewImageFiller(parser *parser.ImageParser) *ImageFiller {
	return &ImageFiller{parser: parser}
}

func (f *ImageFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	image := f.parser.ParseFromURL(container.Short.Image)

	if image == nil {
		return
	}

	service.Image = *image
}
