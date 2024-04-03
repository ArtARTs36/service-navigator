package filler

import (
	"fmt"
	"strings"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
)

const (
	nginxProxyVirtualHostEnv = "VIRTUAL_HOST"
	nginxProxyVirtualPathEnv = "VIRTUAL_PATH"
)

type NginxProxyURLFiller struct {
}

func NewNginxProxyURLFiller() monitor.Filler {
	return &NginxProxyURLFiller{}
}

func (r *NginxProxyURLFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	envVal, envExists := container.Environment[nginxProxyVirtualHostEnv]
	if !envExists || len(envVal) == 0 {
		return
	}

	hosts := strings.Split(envVal, ",")
	if len(hosts) == 0 {
		return
	}

	path := r.getPath(container)

	service.WebURL = fmt.Sprintf("http://%s%s", hosts[0], path)
}

func (r *NginxProxyURLFiller) getPath(container *datastruct.Container) string {
	envVal, envExists := container.Environment[nginxProxyVirtualPathEnv]
	if !envExists || len(envVal) == 0 || envVal == "/" {
		return ""
	}

	if strings.HasPrefix(envVal, "/") {
		return envVal
	}

	return "/" + envVal
}
