package filler

import (
	"github.com/artarts36/service-navigator/internal/services/monitor"
	"github.com/docker/docker/api/types"
)

type CompositeResolver struct {
	fillers []monitor.Filler
}

func NewCompositeFiller(fillers []monitor.Filler) monitor.Filler {
	return &CompositeResolver{
		fillers: fillers,
	}
}

func (r *CompositeResolver) Fill(service *monitor.Service, container *types.ContainerJSON) {
	for _, resolver := range r.fillers {
		resolver.Fill(service, container)
	}
}
