package memory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kashifsoofi/movie-api/internal/store"
)

type MemMoviesStore struct {
	movies []store.Movie
}

func NewMemMoviesStore() store.MoviesStore {
	return &MemMoviesStore{
		movies: []store.Movie{},
	}
}

func (ms *MemMoviesStore) GetAll() ([]store.Movie, error) {
	return ms.movies, nil
}

func (ms *MemMoviesStore) GetByID(ID uuid.UUID) (store.Movie, error) {
	for _, m := range ms.movies {
		if m.ID == ID {
			return m, nil
		}
	}

	return store.Movie{}, errors.New("not found")
}

func (ms *MemMoviesStore) Create(movie store.Movie) {
	ms.movies = append(ms.movies, movie)
}

func (ms *MemMoviesStore) Delete(ID uuid.UUID) (store.Movie, error) {
	index := -1
	for i, m := range ms.movies {
		if m.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return store.Movie{}, errors.New("not found")
	}

	movie := ms.movies[index]
	ms.movies[index] = ms.movies[len(ms.movies)-1]
	ms.movies = ms.movies[:len(ms.movies)-1]
	return movie, nil
}
