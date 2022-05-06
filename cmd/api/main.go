package main

import (
	"context"
	"os"

	"github.com/kashifsoofi/movie-api/internal/api"
	"github.com/kashifsoofi/movie-api/internal/store/memory"
)

func newServer() (*api.Server, error) {
	apiConfig, err := api.LoadConfig()
	if err != nil {
		return nil, err
	}

	store := memory.NewMemoryStore()
	server := api.NewServer(apiConfig.HTTPServer, store)
	return server, nil
}

func main() {
	server, err := newServer()
	if err != nil {
		os.Exit(1)
	}

	ctx := context.Background()
	server.Start(ctx)
}
