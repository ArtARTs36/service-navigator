package weburl

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"strings"
)

const nginxProxyVirtualHostEnv = "VIRTUAL_HOST"

type NginxProxyUrlResolver struct {
}

func (r *NginxProxyUrlResolver) Resolve(container *types.ContainerJSON) string {
	for _, envVar := range container.Config.Env {
		varBag := strings.Split(envVar, "=")

		varName := varBag[0]
		varValue := varBag[1]

		if varName == nginxProxyVirtualHostEnv {
			return fmt.Sprintf("http://%s", varValue)
		}
	}

	return ""
}
