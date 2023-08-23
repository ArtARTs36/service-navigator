package filler

import (
	"github.com/artarts36/service-navigator/internal/service/entity"
)

const labelName = "com.docker.compose.service"

type DCNameFiller struct {
}

func (r *DCNameFiller) Fill(service *entity.Service, container *entity.Container) {
	for key, val := range container.Short.Labels {
		if key == labelName {
			service.Name = val

			return
		}
	}
}
