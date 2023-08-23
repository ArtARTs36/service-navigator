package monitor

import (
	"github.com/docker/docker/api/types"
)

type Filler interface {
	Fill(service *Service, container *types.ContainerJSON)
}
