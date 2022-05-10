package sql

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/kashifsoofi/movie-api/internal/store"
)

const driverName = "pgx"

type SQLStore struct {
	databaseUrl string
	dbx         *sqlx.DB
	movies      SQLMovieStore
}

func NewSQLStore(databaseUrl string) store.Store {
	sqlStore := &SQLStore{
		databaseUrl: databaseUrl,
	}

	sqlStore.movies = NewSQLMovieStore(sqlStore)

	return sqlStore
}

func (s *SQLStore) Connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, s.databaseUrl)
	if err != nil {
		return err
	}

	s.dbx = dbx
	return nil
}

func (s *SQLStore) Close() error {
	return s.dbx.Close()
}

func (s *SQLStore) Movies() store.MovieStore {
	return s.movies
}
