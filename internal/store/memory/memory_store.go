package memory

import "github.com/kashifsoofi/movie-api/internal/store"

type MemoryStore struct {
	movies MemMovieStore
}

func NewMemoryStore() store.Store {
	memoryStore := &MemoryStore{
		movies: NewMemMovieStore(),
	}

	return memoryStore
}

func (ms *MemoryStore) Movies() store.MovieStore {
	return ms.movies
}
