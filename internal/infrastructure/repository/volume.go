package repository

import "github.com/artarts36/service-navigator/internal/domain"

type VolumeRepository struct {
	volumes []*domain.Volume
}

func (r *VolumeRepository) Set(volumes []*domain.Volume) {
	r.volumes = volumes
}

func (r *VolumeRepository) All() []*domain.Volume {
	return r.volumes
}
