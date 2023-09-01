package filler

import (
	"fmt"
	"strings"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
)

const nginxProxyVirtualHostEnv = "VIRTUAL_HOST"

type NginxProxyURLFiller struct {
}

func NewNginxProxyURLFiller() monitor.Filler {
	return &NginxProxyURLFiller{}
}

func (r *NginxProxyURLFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	envVal, envExists := container.Environment[nginxProxyVirtualHostEnv]
	if !envExists {
		return
	}

	if len(envVal) == 0 {
		return
	}

	hosts := strings.Split(envVal, ",")
	if len(hosts) == 0 {
		return
	}

	service.WebURL = fmt.Sprintf("http://%s", hosts[0])
}
