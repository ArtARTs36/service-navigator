package domain

import "github.com/artarts36/service-navigator/internal/shared"

type Service struct {
	Name          string
	WebURL        string
	Status        string
	VCS           *VCS
	ContainerID   string
	Self          bool
	MemoryHistory *shared.MeasurementMetricBuffer
	CPUHistory    *shared.MeasurementMetricBuffer
	Image         ImageShort
}

type ServiceStatus struct {
	Name        string
	WebURL      string
	Status      string
	VCS         *VCS
	ContainerID string
	Self        bool
	Memory      *shared.MeasurementMetric
	CPU         *shared.MeasurementMetric
	Image       ImageShort
}

func NewService(metricDepth int, metricUnique bool) *Service {
	return &Service{
		MemoryHistory: shared.NewMetricBuffer(metricDepth, metricUnique),
		CPUHistory:    shared.NewMetricBuffer(metricDepth, metricUnique),
	}
}
