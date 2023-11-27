package vcs

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/artarts36/service-navigator/internal/domain"
)

const labelGitlabRepository = "org.service_navigator.gitlab_repository"
const labelGithubRepository = "org.service_navigator.github_repository"
const labelBitbucketRepository = "org.service_navigator.bitbucket_repository"

const labelOpenContainerImageSource = "org.opencontainers.image.source"

var NotFoundError = errors.New("vcs not found")

func ParseFromLabels(labels map[string]string) (*domain.VCS, error) {
	var err error

	for key, val := range labels {
		if key == labelGitlabRepository {
			return &domain.VCS{
				Type: "gitlab",
				URL:  val,
			}, nil
		}

		if key == labelGithubRepository {
			return &domain.VCS{
				Type: "github",
				URL:  val,
			}, nil
		}

		if key == labelBitbucketRepository {
			return &domain.VCS{
				Type: "bitbucket",
				URL:  val,
			}, nil
		}

		if key == labelOpenContainerImageSource {
			vcsType, vcsHost, tErr := parseTypeByRawURL(val)

			if tErr != nil {
				err = fmt.Errorf("unable to parse url \"%s\": %v", val, tErr)
			}

			return &domain.VCS{
				Type: vcsType,
				Host: vcsHost,
				URL:  val,
			}, nil
		}
	}

	if err != nil {
		return nil, err
	}

	return nil, NotFoundError
}

func parseTypeByRawURL(rawURL string) (string, string, error) {
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
