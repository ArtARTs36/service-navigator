package config

type Metrics struct {
	Depth      int  `yaml:"depth"`
	OnlyUnique bool `yaml:"only_unique"`
}
