package filler

import (
	"github.com/artarts36/service-navigator/internal/service/entity"
)

const labelGitlabRepository = "org.service_navigator.gitlab_repository"
const labelGithubRepository = "org.service_navigator.github_repository"
const labelBitbucketRepository = "org.service_navigator.bitbucket_repository"

type VCSFiller struct {
}

func (r *VCSFiller) Fill(service *entity.Service, container *entity.Container) {
	for key, val := range container.Short.Labels {
		if key == labelGitlabRepository {
			service.VCS = &entity.VCS{
				Type: "gitlab",
				URL:  val,
			}

			return
		}

		if key == labelGithubRepository {
			service.VCS = &entity.VCS{
				Type: "github",
				URL:  val,
			}

			return
		}

		if key == labelBitbucketRepository {
			service.VCS = &entity.VCS{
				Type: "bitbucket",
				URL:  val,
			}

			return
		}
	}
}
