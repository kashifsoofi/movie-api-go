package memory

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kashifsoofi/movie-api/internal/store"
)

type MemMovieStore struct {
	movies map[uuid.UUID]*store.Movie
}

func NewMemMovieStore() MemMovieStore {
	return MemMovieStore{
		movies: map[uuid.UUID]*store.Movie{},
	}
}

func (s MemMovieStore) GetAll(ctx context.Context) ([]*store.Movie, error) {
	movies := make([]*store.Movie, 0, 0)
	for _, m := range s.movies {
		movies = append(movies, m)
	}
	return movies, nil
}

func (s MemMovieStore) GetByID(ctx context.Context, ID uuid.UUID) (*store.Movie, error) {
	m, ok := s.movies[ID]
	if !ok {
		return nil, errors.New("not found")
	}

	return m, nil
}

func (ms MemMovieStore) Create(ctx context.Context, movie *store.Movie) (*store.Movie, error) {
	movie.ID = uuid.New()
	movie.CreatedAt = time.Now().UTC()
	movie.UpdatedAt = time.Now().UTC()

	ms.movies[movie.ID] = movie
	return movie, nil
}

func (s MemMovieStore) Delete(ctx context.Context, ID uuid.UUID) (*store.Movie, error) {
	m, ok := s.movies[ID]
	if !ok {
		return nil, errors.New("not found")
	}

	delete(s.movies, ID)
	return m, nil
}
