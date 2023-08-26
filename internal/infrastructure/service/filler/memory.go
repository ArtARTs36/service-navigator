package filler

import (
	"time"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/shared"
)

type MemoryFiller struct {
}

func (r *MemoryFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	usedMemory := container.Stats.GetUsedMemory()

	service.Memory = &shared.MeasurementMetric{
		Used:      usedMemory,
		UsedText:  shared.BytesToReadableText(usedMemory),
		Total:     container.Stats.Memory.Limit,
		TotalText: shared.BytesToReadableText(container.Stats.Memory.Limit),
		CreatedAt: time.Now(),
	}
}
