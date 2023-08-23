package entity

import "github.com/docker/docker/api/types"

type Container struct {
	Short *types.Container
	Full  *types.ContainerJSON
}
