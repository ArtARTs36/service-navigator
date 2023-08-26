package datastruct

import "github.com/docker/docker/api/types"

const fullPercent = 100.0

type Stats struct {
	Memory struct {
		Usage int64 `json:"usage"`
		Limit int64 `json:"limit"`
		Stats struct {
			Cache int64 `json:"cache"`
		} `json:"stats"`
	} `json:"memory_stats"`
	CPUStats    types.CPUStats `json:"cpu_stats"`
	PreCPUStats types.CPUStats `json:"precpu_stats"`
}

func (s *Stats) GetUsedMemory() int64 {
	return s.Memory.Usage - s.Memory.Stats.Cache
}

func (s *Stats) GetCPUUsage() float64 {
	cpuPercent := 0.0
	cpuDelta := float64(s.CPUStats.CPUUsage.TotalUsage) - float64(s.PreCPUStats.SystemUsage)
	sysDelta := float64(s.CPUStats.SystemUsage) - float64(s.PreCPUStats.SystemUsage)

	if sysDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / sysDelta) * float64(len(s.CPUStats.CPUUsage.PercpuUsage)) * fullPercent
	}

	return cpuPercent
}
