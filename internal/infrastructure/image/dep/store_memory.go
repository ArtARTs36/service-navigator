package dep

import (
	"context"

	"github.com/artarts36/depexplorer"
)

type MemoryFileStore struct {
	cache map[string]map[depexplorer.DependencyManager]*depexplorer.File
}

func NewMemoryFileStore() *MemoryFileStore {
	return &MemoryFileStore{cache: make(map[string]map[depexplorer.DependencyManager]*depexplorer.File)}
}

func (s *MemoryFileStore) Get(
	_ context.Context,
	image string,
) (map[depexplorer.DependencyManager]*depexplorer.File, error) {
	dfile, ok := s.cache[image]
	if !ok {
		return nil, ErrFilesNotFound
	}

	return dfile, ErrFilesNotFound
}

func (s *MemoryFileStore) Set(
	_ context.Context,
	image string,
	dfile map[depexplorer.DependencyManager]*depexplorer.File,
) error {
	s.cache[image] = dfile
	return nil
}
