package parser

import (
	"testing"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestImageParser_ParseFromURI(t *testing.T) {
	cases := []struct {
		URI      string
		Expected *domain.ImageShort
	}{
		{
			URI: "php:8.0",
			Expected: &domain.ImageShort{
				Name:        "php",
				Version:     "8.0",
				RegistryURL: "https://hub.docker.com/_/php",
			},
		},
		{
			URI: "bitnami/kafka",
			Expected: &domain.ImageShort{
				Name:        "bitnami/kafka",
				Version:     "latest",
				RegistryURL: "https://hub.docker.com/r/bitnami/kafka",
			},
		},
		{
			URI: "bitnami/kafka:1.2.3",
			Expected: &domain.ImageShort{
				Name:        "bitnami/kafka",
				Version:     "1.2.3",
				RegistryURL: "https://hub.docker.com/r/bitnami/kafka",
			},
		},
		{
			URI: "ghcr.io/home-assistant/home-assistant:stable",
			Expected: &domain.ImageShort{
				Name:        "ghcr.io/home-assistant/home-assistant",
				Version:     "stable",
				RegistryURL: "http://ghcr.io/home-assistant/home-assistant",
			},
		},
	}

	p := &ImageParser{}

	for _, tCase := range cases {
		givenImage := p.ParseFromURL(tCase.URI)

		assert.Equal(t, tCase.Expected, givenImage)
	}
}
