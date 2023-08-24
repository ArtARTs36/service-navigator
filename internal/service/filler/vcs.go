package filler

import (
	"fmt"
	"log"
	"net/url"

	"github.com/artarts36/service-navigator/internal/service/entity"
)

const labelGitlabRepository = "org.service_navigator.gitlab_repository"
const labelGithubRepository = "org.service_navigator.github_repository"
const labelBitbucketRepository = "org.service_navigator.bitbucket_repository"

const labelOpenContainerImageSource = "org.opencontainers.image.source"

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

		if key == labelOpenContainerImageSource {
			vcsType, err := r.resolveTypeByRawURL(val)

			if err != nil {
				log.Print(err)

				continue
			}

			service.VCS = &entity.VCS{
				Type: vcsType,
				URL:  val,
			}

			return
		}
	}
}

func (r *VCSFiller) resolveTypeByRawURL(rawURL string) (string, error) {
	vcsURL, err := url.Parse(rawURL)

	if err != nil {
		return "", fmt.Errorf("unable to parse url \"%s\": %w", rawURL, err)
	}

	switch vcsURL.Host {
	case "github.com":
		return "github", nil
	case "gitlab.com":
		return "gitlab", nil
	case "bitbucket.com":
		return "bitbucket", nil
	}

	return "", nil
}
