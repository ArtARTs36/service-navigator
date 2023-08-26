package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/service/entity"
	"github.com/artarts36/service-navigator/internal/service/monitor"
)

type CompositeFiller struct {
	fillers []monitor.Filler
}

func NewCompositeFiller(fillers []monitor.Filler) monitor.Filler {
	return &CompositeFiller{
		fillers: fillers,
	}
}

func (r *CompositeFiller) Fill(service *domain.ServiceStatus, container *entity.Container) {
	for _, resolver := range r.fillers {
		resolver.Fill(service, container)
	}
}
