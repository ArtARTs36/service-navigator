package filler

import (
	"github.com/artarts36/service-navigator/internal/service/entity"
)

type VCSFiller struct {
}

func (r *VCSFiller) Fill(service *entity.Service, container *entity.Container) {
	for key, val := range container.Short.Labels {
		if key == "service_navigator.gitlab_repository" {
			service.VCS = &entity.VCS{
				Type: "gitlab",
				URL:  val,
			}

			return
		}

		if key == "service_navigator.github_repository" {
			service.VCS = &entity.VCS{
				Type: "github",
				URL:  val,
			}

			return
		}

		if key == "service_navigator.bitbucket_repository" {
			service.VCS = &entity.VCS{
				Type: "bitbucket",
				URL:  val,
			}

			return
		}
	}
}
