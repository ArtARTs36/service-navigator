package filler

import (
	"fmt"
	"time"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/shared"
)

const fullPercent = 100

type CPUFiller struct {
}

func (r *CPUFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	cpuUsage := container.Stats.GetCPUUsage()

	service.CPU = &shared.MeasurementMetric{
		Used:      int64(cpuUsage),
		UsedText:  fmt.Sprintf("%.2f", cpuUsage) + "%",
		Total:     fullPercent,
		TotalText: "100%",
		CreatedAt: time.Now(),
	}
}
