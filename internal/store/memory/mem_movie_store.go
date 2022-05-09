package memory

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kashifsoofi/movie-api/internal/store"
)

type MemMovieStore struct {
	movies []*store.Movie
}

func NewMemMovieStore() MemMovieStore {
	return MemMovieStore{
		movies: []*store.Movie{},
	}
}

func (s MemMovieStore) GetAll(ctx context.Context) ([]*store.Movie, error) {
	return s.movies, nil
}

func (s MemMovieStore) GetByID(ctx context.Context, ID uuid.UUID) (*store.Movie, error) {
	for _, m := range s.movies {
		if m.ID == ID {
			return m, nil
		}
	}

	return nil, errors.New("not found")
}

func (ms MemMovieStore) Create(ctx context.Context, movie *store.Movie) (*store.Movie, error) {
	movie.ID = uuid.New()
	movie.CreatedAt = time.Now().UTC()
	movie.UpdatedAt = time.Now().UTC()

	ms.movies = append(ms.movies, movie)
	return movie, nil
}

func (s MemMovieStore) Delete(ctx context.Context, ID uuid.UUID) (*store.Movie, error) {
	index := -1
	for i, m := range s.movies {
		if m.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, errors.New("not found")
	}

	movie := s.movies[index]
	s.movies[index] = s.movies[len(s.movies)-1]
	s.movies = s.movies[:len(s.movies)-1]
	return movie, nil
}
