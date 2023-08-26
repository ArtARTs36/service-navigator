package filler

import (
	"testing"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
)

func TestImageFiller_Fill(t *testing.T) {
	cases := []struct {
		containerShortImage string
		expectedImage       domain.Image
	}{
		{
			containerShortImage: "vendor/image",
			expectedImage: domain.Image{
				Name:        "vendor/image",
				Version:     "latest",
				RegistryURL: "https://hub.docker.com/r/vendor/image",
			},
		},
		{
			containerShortImage: "vendor/image:v1",
			expectedImage: domain.Image{
				Name:        "vendor/image",
				Version:     "v1",
				RegistryURL: "https://hub.docker.com/r/vendor/image",
			},
		},
		{
			containerShortImage: "image:v1",
			expectedImage: domain.Image{
				Name:        "image",
				Version:     "v1",
				RegistryURL: "https://hub.docker.com/_/image",
			},
		},
		{
			containerShortImage: "registry.io/vendor/image",
			expectedImage: domain.Image{
				Name:        "registry.io/vendor/image",
				Version:     "latest",
				RegistryURL: "http://registry.io/vendor/image",
			},
		},
		{
			containerShortImage: "registry.io/vendor/image:v1",
			expectedImage: domain.Image{
				Name:        "registry.io/vendor/image",
				Version:     "v1",
				RegistryURL: "http://registry.io/vendor/image",
			},
		},
	}

	filler := ImageFiller{}

	for _, tCase := range cases {
		service := &domain.ServiceStatus{}

		cont := &datastruct.Container{
			Short: types.Container{
				Image: tCase.containerShortImage,
			},
		}

		filler.Fill(service, cont)

		assert.Equal(t, tCase.expectedImage, service.Image)
	}
}
