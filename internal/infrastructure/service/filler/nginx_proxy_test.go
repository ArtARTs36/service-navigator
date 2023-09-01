package filler_test

import (
	"testing"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/stretchr/testify/assert"
)

func TestNginxProxyFiller_Fill(t *testing.T) {
	cases := []struct {
		envVarValue    string
		expectedWebURL string
	}{
		{
			envVarValue:    "",
			expectedWebURL: "",
		},
		{
			envVarValue:    "domain.com",
			expectedWebURL: "http://domain.com",
		},
		{
			envVarValue:    "domain.com,site.com",
			expectedWebURL: "http://domain.com",
		},
	}

	imgFiller := filler.NginxProxyURLFiller{}

	for _, tCase := range cases {
		service := &domain.ServiceStatus{}

		cont := &datastruct.Container{
			Environment: map[string]string{
				"VIRTUAL_HOST": tCase.envVarValue,
			},
		}

		imgFiller.Fill(service, cont)

		assert.Equal(t, tCase.expectedWebURL, service.WebURL)
	}
}
