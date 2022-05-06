package store

import (
	"time"

	"github.com/google/uuid"
)

type Store interface {
	Movies() MoviesStore
}

type Movie struct {
	ID          uuid.UUID
	Title       string
	Director    string
	ReleaseDate time.Time
	TicketPrice float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MoviesStore interface {
	GetAll() ([]Movie, error)
	GetByID(uuid.UUID) (Movie, error)
	Create(Movie)
	Delete(uuid.UUID) (Movie, error)
}
