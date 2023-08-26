package monitor

import (
	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/service/entity"
)

type Filler interface {
	Fill(service *domain.ServiceStatus, container *entity.Container)
}
