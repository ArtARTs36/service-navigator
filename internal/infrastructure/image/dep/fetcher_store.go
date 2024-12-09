package dep

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/artarts36/depexplorer"

	"github.com/artarts36/service-navigator/internal/domain"
)

type StoreableFetcher struct {
	store FileStore
	inner Fetcher
}

func NewStoreableFetcher(inner Fetcher, store FileStore) *StoreableFetcher {
	return &StoreableFetcher{
		store: store,
		inner: inner,
	}
}

func (f *StoreableFetcher) Fetch(
	ctx context.Context,
	image *domain.Image,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	key := f.key(image)

	file, err := f.store.Get(ctx, key)
	if err != nil {
		if errors.Is(err, ErrFilesNotFound) {
			return f.fetchAndStore(ctx, image, key)
		}
		return nil, err
	}

	return file, nil
}

func (f *StoreableFetcher) fetchAndStore(
	ctx context.Context,
	image *domain.Image,
	key string,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	file, err := f.inner.Fetch(ctx, image)
	if err != nil {
		return nil, err
	}

	err = f.store.Set(ctx, key, file)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to store image: %v", err)
	}

	return file, nil
}

func (*StoreableFetcher) key(img *domain.Image) string {
	if img.NameDetails.IsLatest() {
		return img.ID
	}
	return img.Name
}
