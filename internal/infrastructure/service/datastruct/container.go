package datastruct

import "github.com/docker/docker/api/types"

type Container struct {
	Short types.Container
	Full  types.ContainerJSON
	Stats *Stats `json:"stats"`
}
