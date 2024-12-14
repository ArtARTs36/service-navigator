package filler

import (
	"github.com/artarts36/depexplorer"
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type LanguagesFiller struct {
}

func NewLanguageFiller() *LanguagesFiller {
	return &LanguagesFiller{}
}

func (f *LanguagesFiller) Fill(image *domain.Image, _ *datastruct.ImageMeta) {
	if image.DepFiles == nil {
		return
	}

	langs := make([]*domain.Language, 0, len(image.DepFiles))
	frameworks := make([]*depexplorer.Framework, 0, len(image.DepFiles))

	for _, file := range image.DepFiles {
		lang := &domain.Language{
			Name: string(file.Language.Name),
		}
		if file.Language.Version != nil {
			lang.Version = file.Language.Version.Full
		}
		langs = append(langs, lang)

		frameworks = append(frameworks, file.Frameworks...)
	}

	image.Languages = langs
	image.Frameworks = frameworks
}
