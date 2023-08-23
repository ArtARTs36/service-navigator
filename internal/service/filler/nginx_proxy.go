package filler

import (
	"fmt"
	"github.com/artarts36/service-navigator/internal/service/entity"
	"strings"
)

const nginxProxyVirtualHostEnv = "VIRTUAL_HOST"

type NginxProxyUrlFiller struct {
}

func (r *NginxProxyUrlFiller) Fill(service *entity.Service, container *entity.Container) {
	for _, envVar := range container.Full.Config.Env {
		varBag := strings.Split(envVar, "=")

		varName := varBag[0]
		varValue := varBag[1]

		if varName == nginxProxyVirtualHostEnv {
			service.WebUrl = fmt.Sprintf("http://%s", varValue)

			return
		}
	}
}
