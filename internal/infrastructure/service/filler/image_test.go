package filler_test

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/image/parser"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
)

func TestImageFiller_Fill(t *testing.T) {
	cases := []struct {
		containerShortImage string
		expectedImage       domain.ImageShort
	}{
		{
			containerShortImage: "vendor/image",
			expectedImage: domain.ImageShort{
				Name:        "vendor/image",
				Version:     "latest",
				RegistryURL: "https://hub.docker.com/r/vendor/image",
			},
		},
		{
			containerShortImage: "vendor/image:v1",
			expectedImage: domain.ImageShort{
				Name:        "vendor/image",
				Version:     "v1",
				RegistryURL: "https://hub.docker.com/r/vendor/image",
			},
		},
		{
			containerShortImage: "image:v1",
			expectedImage: domain.ImageShort{
				Name:        "image",
				Version:     "v1",
				RegistryURL: "https://hub.docker.com/_/image",
			},
		},
		{
			containerShortImage: "registry.io/vendor/image",
			expectedImage: domain.ImageShort{
				Name:        "registry.io/vendor/image",
				Version:     "latest",
				RegistryURL: "http://registry.io/vendor/image",
			},
		},
		{
			containerShortImage: "registry.io/vendor/image:v1",
			expectedImage: domain.ImageShort{
				Name:        "registry.io/vendor/image",
				Version:     "v1",
				RegistryURL: "http://registry.io/vendor/image",
			},
		},
	}

	imgFiller := filler.NewImageFiller(&parser.ImageParser{})

	for _, tCase := range cases {
		service := &domain.ServiceStatus{}

		cont := &datastruct.Container{
			Short: types.Container{
				Image: tCase.containerShortImage,
			},
		}

		imgFiller.Fill(service, cont)

		assert.Equal(t, tCase.expectedImage, service.Image)
	}
}
