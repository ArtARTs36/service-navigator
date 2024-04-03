package filler_test

import (
	"fmt"
	"testing"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/stretchr/testify/assert"
)

func TestNginxProxyFiller_Fill(t *testing.T) {
	cases := []struct {
		env            map[string]string
		expectedWebURL string
	}{
		{
			env:            map[string]string{},
			expectedWebURL: "",
		},
		{
			env: map[string]string{
				"VIRTUAL_HOST": "domain.com",
			},
			expectedWebURL: "http://domain.com",
		},
		{
			env: map[string]string{
				"VIRTUAL_HOST": "domain.com,site.com",
			},
			expectedWebURL: "http://domain.com",
		},
		{
			env: map[string]string{
				"VIRTUAL_HOST": "domain.com,site.com",
				"VIRTUAL_PATH": "/api",
			},
			expectedWebURL: "http://domain.com/api",
		},
		{
			env: map[string]string{
				"VIRTUAL_HOST": "domain.com,site.com",
				"VIRTUAL_PATH": "api",
			},
			expectedWebURL: "http://domain.com/api",
		},
	}

	imgFiller := filler.NginxProxyURLFiller{}

	for i, tCase := range cases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			service := &domain.ServiceStatus{}

			cont := &datastruct.Container{
				Environment: tCase.env,
			}

			imgFiller.Fill(service, cont)

			assert.Equal(t, tCase.expectedWebURL, service.WebURL)
		})
	}
}
