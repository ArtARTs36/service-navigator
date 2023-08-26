package domain

import "github.com/artarts36/service-navigator/internal/shared"

type Service struct {
	Name          string
	WebURL        string
	Status        string
	VCS           *VCS
	ContainerID   string
	Self          bool
	MemoryHistory *shared.MetricBuffer
}

type ServiceStatus struct {
	Name        string
	WebURL      string
	Status      string
	VCS         *VCS
	ContainerID string
	Self        bool
	Memory      *shared.Metric
}

func NewService(metricDepth int) *Service {
	return &Service{
		MemoryHistory: shared.NewMetricBuffer(metricDepth),
	}
}
