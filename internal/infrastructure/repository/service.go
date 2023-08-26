package repository

import (
	"github.com/artarts36/service-navigator/internal/domain"
)

type ServiceRepository struct {
	services []*domain.Service
}

func (r *ServiceRepository) Set(services []*domain.Service) {
	r.services = services
}

func (r *ServiceRepository) All() []*domain.Service {
	return r.services
}
