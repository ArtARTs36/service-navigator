package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

const labelName = "com.docker.compose.service"

type DCNameFiller struct {
}

func (r *DCNameFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	for key, val := range container.Short.Labels {
		if key == labelName {
			service.Name = val

			return
		}
	}
}
