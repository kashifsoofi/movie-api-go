package memory

import "github.com/kashifsoofi/movie-api/internal/store"

type MemoryStore struct {
	movies store.MoviesStore
}

func NewMemoryStore() store.Store {
	memoryStore := &MemoryStore{
		movies: NewMemMoviesStore(),
	}

	return memoryStore
}

func (ms *MemoryStore) Movies() store.MoviesStore {
	return ms.movies
}
