package filler

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
	"github.com/mitchellh/hashstructure"
)

type CompositeFiller struct {
	fillers []monitor.Filler
}

func NewCompositeFiller(fillers []monitor.Filler) monitor.Filler {
	return &CompositeFiller{
		fillers: fillers,
	}
}

func (r *CompositeFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	for _, resolver := range r.fillers {
		resolver.Fill(service, container)
	}
}

type OrFiller struct {
	fillers []monitor.Filler
}

func NewOrFiller(fillers []monitor.Filler) monitor.Filler {
	return &OrFiller{
		fillers: fillers,
	}
}

func (r *OrFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	prevHash, _ := hashstructure.Hash(service, nil)

	for _, resolver := range r.fillers {
		resolver.Fill(service, container)

		newHash, _ := hashstructure.Hash(service, nil)

		if prevHash != newHash {
			return
		}
	}
}
