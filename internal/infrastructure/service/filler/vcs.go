package filler

import (
	"fmt"
	"log"
	"net/url"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

const labelGitlabRepository = "org.service_navigator.gitlab_repository"
const labelGithubRepository = "org.service_navigator.github_repository"
const labelBitbucketRepository = "org.service_navigator.bitbucket_repository"

const labelOpenContainerImageSource = "org.opencontainers.image.source"

type VCSFiller struct {
}

func (r *VCSFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	for key, val := range container.Short.Labels {
		if key == labelGitlabRepository {
			service.VCS = &domain.VCS{
				Type: "gitlab",
				URL:  val,
			}

			return
		}

		if key == labelGithubRepository {
			service.VCS = &domain.VCS{
				Type: "github",
				URL:  val,
			}

			return
		}

		if key == labelBitbucketRepository {
			service.VCS = &domain.VCS{
				Type: "bitbucket",
				URL:  val,
			}

			return
		}

		if key == labelOpenContainerImageSource {
			vcsType, vcsHost, err := r.resolveTypeByRawURL(val)

			if err != nil {
				log.Printf("unable to parse url \"%s\": %v", val, err)

				continue
			}

			service.VCS = &domain.VCS{
				Type: vcsType,
				Host: vcsHost,
				URL:  val,
			}

			return
		}
	}
}

func (r *VCSFiller) resolveTypeByRawURL(rawURL string) (string, string, error) {
	vcsURL, err := url.Parse(rawURL)

	if err != nil {
		return "", "", fmt.Errorf("unable to parse url \"%s\": %w", rawURL, err)
	}

	switch vcsURL.Host {
	case "github.com":
		return "github", vcsURL.Host, nil
	case "gitlab.com":
		return "gitlab", vcsURL.Host, nil
	case "bitbucket.com":
		return "bitbucket", vcsURL.Host, nil
	}

	return "", vcsURL.Host, nil
}