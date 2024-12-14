package dep

import (
	"context"
	"errors"
	"strings"

	"github.com/artarts36/depexplorer/pkg/repository"

	"github.com/artarts36/depexplorer"
	"github.com/artarts36/service-navigator/internal/domain"
)

var ErrVCSNotSupported = errors.New("VCS not supported")

type Fetcher interface {
	Fetch(ctx context.Context, image *domain.Image) (map[depexplorer.DependencyManager]*depexplorer.File, error)
}

type ClientFetcher struct {
	explorers map[string]repository.Explorer
}

func NewClientFetcher(
	explorers map[string]repository.Explorer,
) *ClientFetcher {
	return &ClientFetcher{
		explorers: explorers,
	}
}

func (f *ClientFetcher) Fetch(
	ctx context.Context,
	image *domain.Image,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	repoURLParts := strings.Split(image.VCS.URL, "/")

	explorer, ok := f.explorers[image.VCS.Type]
	if !ok {
		return nil, ErrVCSNotSupported
	}

	depFile, err := explorer.ExploreRepository(
		ctx,
		repository.Repo{
			Owner: repoURLParts[len(repoURLParts)-2],
			Name:  repoURLParts[len(repoURLParts)-1],
		},
	)
	if err != nil {
		return nil, err
	}

	return depFile, nil
}
