package memory

import (
	"context"

	"github.com/kashifsoofi/movie-api/internal/store"
)

type MemoryStore struct {
	movies MemMovieStore
}

func NewMemoryStore() store.Store {
	memoryStore := &MemoryStore{
		movies: NewMemMovieStore(),
	}

	return memoryStore
}

func (s *MemoryStore) Connect(ctx context.Context) error {
	return nil
}

func (s *MemoryStore) Close() error {
	return nil
}

func (s *MemoryStore) Movies() store.MovieStore {
	return s.movies
}
