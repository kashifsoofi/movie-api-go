package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Store interface {
	Movies() MovieStore
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

type MovieStore interface {
	GetAll(context.Context) ([]*Movie, error)
	GetByID(context.Context, uuid.UUID) (*Movie, error)
	Create(context.Context, *Movie) (*Movie, error)
	Delete(context.Context, uuid.UUID) (*Movie, error)
}
