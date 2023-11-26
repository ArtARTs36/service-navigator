package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/parser"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type ShortFiller struct {
	parser *parser.ImageParser
}

func NewShortFiller(parser *parser.ImageParser) *ShortFiller {
	return &ShortFiller{parser: parser}
}

func (f *ShortFiller) Fill(image *domain.Image, meta *datastruct.ImageMeta) {
	short := f.parser.ParseFromURL(meta.URI)

	if short == nil {
		return
	}

	image.Short = *short
}
