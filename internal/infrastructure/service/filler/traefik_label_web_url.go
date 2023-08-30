package filler

import (
	"fmt"
	"regexp"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

type TraefikLabelWebUrlFiller struct {
}

func (r *TraefikLabelWebUrlFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	var re = regexp.MustCompile(`(?m)traefik\.http\.routers\..*\.rule`)

	for lKey, val := range container.Short.Labels {
		if re.FindString(lKey) == "" {
			continue
		}

		valRe := regexp.MustCompile("(?m)Host\\(`(.*)`\\)")

		matches := valRe.FindStringSubmatch(val)

		if len(matches) != 2 {
			continue
		}

		service.WebURL = fmt.Sprintf("http://%s", matches[1])
	}
}
