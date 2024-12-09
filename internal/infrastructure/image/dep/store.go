package dep

import (
	"context"
	"errors"

	"github.com/artarts36/depexplorer"
)

var ErrFilesNotFound = errors.New("files not found")

type FileStore interface {
	Get(ctx context.Context, image string) (map[depexplorer.DependencyManager]*depexplorer.File, error)
	Set(ctx context.Context, image string, file map[depexplorer.DependencyManager]*depexplorer.File) error
}
