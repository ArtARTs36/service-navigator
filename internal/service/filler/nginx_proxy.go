package filler

import (
	"fmt"
	"strings"

	"github.com/artarts36/service-navigator/internal/service/entity"
)

const nginxProxyVirtualHostEnv = "VIRTUAL_HOST"

type NginxProxyURLFiller struct {
}

func (r *NginxProxyURLFiller) Fill(service *entity.Service, container *entity.Container) {
	for _, envVar := range container.Full.Config.Env {
		varBag := strings.Split(envVar, "=")

		varName := varBag[0]
		varValue := varBag[1]

		if varName == nginxProxyVirtualHostEnv {
			service.WebURL = fmt.Sprintf("http://%s", varValue)

			return
		}
	}
}
