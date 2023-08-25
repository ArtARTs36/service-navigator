package repository

import "github.com/artarts36/service-navigator/internal/service/entity"

type ServiceRepository struct {
	services []*entity.Service
}

func (r *ServiceRepository) Set(services []*entity.Service) {
	r.services = services
}

func (r *ServiceRepository) All() []*entity.Service {
	return r.services
}
