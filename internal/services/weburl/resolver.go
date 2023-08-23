package weburl

import "github.com/docker/docker/api/types"

type UrlResolver interface {
	Resolve(container *types.ContainerJSON) string
}
