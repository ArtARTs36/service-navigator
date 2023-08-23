package filler

import (
	"fmt"
	"github.com/artarts36/service-navigator/internal/services/monitor"
	"github.com/docker/docker/api/types"
	"strings"
)

const nginxProxyVirtualHostEnv = "VIRTUAL_HOST"

type NginxProxyUrlFiller struct {
}

func (r *NginxProxyUrlFiller) Fill(service *monitor.Service, container *types.ContainerJSON) {
	for _, envVar := range container.Config.Env {
		varBag := strings.Split(envVar, "=")

		varName := varBag[0]
		varValue := varBag[1]

		if varName == nginxProxyVirtualHostEnv {
			service.WebUrl = fmt.Sprintf("http://%s", varValue)

			return
		}
	}
}
