package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

const labelName = "com.docker.compose.service"

type DCNameFiller struct {
}

func (r *DCNameFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	label, labelExists := container.Short.Labels[labelName]

	if labelExists {
		service.Name = label
	}
}
