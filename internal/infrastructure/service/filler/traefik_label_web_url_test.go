package filler_test

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTraefikLabelWebUrlFiller(t *testing.T) {
	cases := []struct {
		Labels         map[string]string
		ExpectedWebURL string
	}{
		{
			Labels:         map[string]string{},
			ExpectedWebURL: "",
		},
		{
			Labels: map[string]string{
				"other-label": "",
			},
			ExpectedWebURL: "",
		},
		{
			Labels: map[string]string{
				"other-label": "Host(`site.com`)",
			},
			ExpectedWebURL: "",
		},
		{
			Labels: map[string]string{
				"traefik.http.routers.domain-traefik.rule": "other-rule",
			},
			ExpectedWebURL: "",
		},
		{
			Labels: map[string]string{
				"traefik.http.routers.domain-traefik.rule": "Host(`site.com`)",
			},
			ExpectedWebURL: "http://site.com",
		},
	}

	f := filler.TraefikLabelWebUrlFiller{}

	for _, tCase := range cases {
		service := &domain.ServiceStatus{}

		f.Fill(service, &datastruct.Container{
			Short: types.Container{
				Labels: tCase.Labels,
			},
		})

		assert.Equal(t, tCase.ExpectedWebURL, service.WebURL)
	}
}
