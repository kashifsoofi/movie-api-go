package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kashifsoofi/movie-api/internal/store"
)

type SQLMovieStore struct {
	*SQLStore
}

func NewSQLMovieStore(sqlStore *SQLStore) SQLMovieStore {
	return SQLMovieStore{sqlStore}
}

func (s SQLMovieStore) GetAll(ctx context.Context) ([]*store.Movie, error) {
	var movies []*store.Movie
	if err := s.dbx.SelectContext(
		ctx,
		&movies,
		`SELECT
			id, title, director, release_date, ticket_price, created_at, updated_at
		FROM movies`); err != nil {
		return nil, fmt.Errorf("could not list movies, err: %w", err)
	}
	return movies, nil
}

func (s SQLMovieStore) GetByID(ctx context.Context, ID uuid.UUID) (*store.Movie, error) {
	var movie store.Movie
	if err := s.dbx.GetContext(
		ctx,
		&movie,
		`SELECT
			id, title, director, release_date, ticket_price, created_at, updated_at
		FROM movies
		WHERE id = ?`,
		ID); err != nil {
		if err != sql.ErrNoRows {

		}

		return nil, errors.New("not found")
	}

	return &movie, nil
}

func (s SQLMovieStore) Create(ctx context.Context, movie *store.Movie) (*store.Movie, error) {
	movie.ID = uuid.New()
	if _, err := s.dbx.NamedExecContext(
		ctx,
		`INSERT INTO movies
			(id, title, director, release_date, ticket_price)
		VALUES
			(:id, :title, :director, :relase_date, :ticket_price)`,
		movie); err != nil {
		return nil, fmt.Errorf("count not insert movie, err: %w", err)
	}

	return s.GetByID(ctx, movie.ID)
}

func (s SQLMovieStore) Delete(ctx context.Context, ID uuid.UUID) (*store.Movie, error) {
	movie, err := s.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	if _, err := s.dbx.NamedExecContext(
		ctx,
		`DELETE FROM movies
		WHERE id = :id`, ID); err != nil {
		return nil, fmt.Errorf("count not delete movie, err: %w", err)
	}

	return movie, nil
}
