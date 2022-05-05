package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUrl string `required:"true" split_words:"true"`
	HostPort    int    `default:"8080"`
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

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
}
