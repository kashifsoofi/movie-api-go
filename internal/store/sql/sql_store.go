package sql

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kashifsoofi/movie-api/internal/store"
	_ "github.com/lib/pq"
)

const driverName = "postgres"

type SQLStore struct {
	dbx    *sqlx.DB
	movies SQLMovieStore
}

func NewSQLStore(databaseUrl string) store.Store {
	dbx, err := sqlx.Connect(driverName, databaseUrl)
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
