package filler

import (
	"fmt"
	"strings"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

const nginxProxyVirtualHostEnv = "VIRTUAL_HOST"

type NginxProxyURLFiller struct {
}

func (r *NginxProxyURLFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
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
