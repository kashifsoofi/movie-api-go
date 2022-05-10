package sql

import (
	"context"
	"log"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/kashifsoofi/movie-api/internal/store"
)

const driverName = "postgres"

type SQLStore struct {
	dbx    *sqlx.DB
	movies SQLMovieStore
}

func NewSQLStore(ctx context.Context, databaseUrl string) store.Store {
	dbx, err := sqlx.ConnectContext(ctx, driverName, databaseUrl)
	if err != nil {
		log.Fatalln(err)
	}

	sqlStore := &SQLStore{
		dbx: dbx,
	}

	sqlStore.movies = NewSQLMovieStore(sqlStore)

	return sqlStore
}

func (ms *SQLStore) Movies() store.MovieStore {
	return ms.movies
}
