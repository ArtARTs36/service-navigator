package dep

import (
	"context"
	"errors"
	"strings"

	"github.com/artarts36/depexplorer"
	githubExplorer "github.com/artarts36/depexplorer/pkg/github"
	"github.com/artarts36/service-navigator/internal/domain"
	githubClient "github.com/google/go-github/v67/github"
	log "github.com/sirupsen/logrus"
)

var ErrVCSNotSupported = errors.New("VCS not supported")

type Fetcher interface {
	Fetch(ctx context.Context, image *domain.Image) (map[depexplorer.DependencyManager]*depexplorer.File, error)
}

type ClientFetcher struct {
	githubClient *githubClient.Client
}

func depexplorerLogger() githubExplorer.Logger {
	return func(s string, m map[string]interface{}) {
		log.Info(s, m)
	}
}

func NewClientFetcher(
	githubClient *githubClient.Client,
) *ClientFetcher {
	return &ClientFetcher{
		githubClient: githubClient,
	}
}

func (f *ClientFetcher) Fetch(
	ctx context.Context,
	image *domain.Image,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	switch image.VCS.Type {
	case "github":
		return f.fetchGithub(ctx, image)
	default:
		return nil, ErrVCSNotSupported
	}
}

func (f *ClientFetcher) fetchGithub(
	ctx context.Context,
	image *domain.Image,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	repoURLParts := strings.Split(image.VCS.URL, "/")

	depFile, err := githubExplorer.ScanRepository(
		ctx,
		f.githubClient,
		githubExplorer.Repository{
			Owner: repoURLParts[len(repoURLParts)-2],
			Repo:  repoURLParts[len(repoURLParts)-1],
		},
		depexplorerLogger(),
	)
	if err != nil {
		return nil, err
	}
	return depFile, nil
}
