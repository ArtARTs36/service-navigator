package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

const swarmServiceNameLabel = "com.docker.swarm.service.name"

type SwarmNameFiller struct {
}

func (r *SwarmNameFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	label, labelExists := container.Short.Labels[swarmServiceNameLabel]
	if !labelExists {
		return
	}

	service.Name = label
}
