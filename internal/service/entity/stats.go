package entity

type Stats struct {
	Memory struct {
		Usage int64 `json:"usage"`
		Limit int64 `json:"limit"`
		Stats struct {
			Cache int64 `json:"cache"`
		} `json:"stats"`
	} `json:"memory_stats"`
}

func (s *Stats) GetUsedMemory() int64 {
	return s.Memory.Usage - s.Memory.Stats.Cache
}
