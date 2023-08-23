package monitor

import (
	"github.com/artarts36/service-navigator/internal/service/entity"
)

type Filler interface {
	Fill(service *entity.Service, container *entity.Container)
}
