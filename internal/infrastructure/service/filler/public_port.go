package filler

import (
	"fmt"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
)

type PublicPortFiller struct {
}

func NewPublicPortFiller() monitor.Filler {
	return &PublicPortFiller{}
}

func (r *PublicPortFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	for _, port := range container.Short.Ports {
		if port.PublicPort != 0 {
			service.WebURL = fmt.Sprintf("http://localhost:%d", port.PublicPort)
		}
	}
}
