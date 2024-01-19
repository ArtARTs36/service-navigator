package parser

import (
	"testing"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestImageParser_ParseFromURI(t *testing.T) {
	cases := []struct {
		URI      string
		Expected *domain.NameDetails
	}{
		{
			URI: "php:8.0",
			Expected: &domain.NameDetails{
				Name:        "php",
				Version:     "8.0",
				RegistryURL: "https://hub.docker.com/_/php",
			},
		},
		{
			URI: "bitnami/kafka",
			Expected: &domain.NameDetails{
				Name:        "kafka",
				Version:     "latest",
				RegistryURL: "https://hub.docker.com/r/bitnami/kafka",
				Vendor:      "bitnami",
			},
		},
		{
			URI: "bitnami/kafka:1.2.3",
			Expected: &domain.NameDetails{
				Name:        "kafka",
				Version:     "1.2.3",
				RegistryURL: "https://hub.docker.com/r/bitnami/kafka",
				Vendor:      "bitnami",
			},
		},
		{
			URI: "ghcr.io/home-assistant/project:stable",
			Expected: &domain.NameDetails{
				Name:        "project",
				Version:     "stable",
				RegistryURL: "http://ghcr.io/home-assistant/project",
				Vendor:      "home-assistant",
			},
		},
	}

	p := &ImageParser{}

	for _, tCase := range cases {
		givenImage := p.ParseFromURL(tCase.URI)

		assert.Equal(t, tCase.Expected, givenImage)
	}
}
