package main

import (
	"context"
	"os"

	"github.com/kashifsoofi/movie-api/internal/api"
	"github.com/kashifsoofi/movie-api/internal/store/memory"
	"github.com/kashifsoofi/movie-api/internal/store/sql"
)

func newServer(ctx context.Context) (*api.Server, error) {
	apiConfig, err := api.LoadConfig()
	if err != nil {
		return nil, err
	}

	store := memory.NewMemoryStore()
	if apiConfig.Store == "sql" {
		store = sql.NewSQLStore(ctx, apiConfig.DatabaseURL)
	}
	server := api.NewServer(apiConfig.HTTPServer, store)
	return server, nil
}

func main() {
	ctx := context.Background()
	server, err := newServer(ctx)
	if err != nil {
		os.Exit(1)
	}

	server.Start(ctx)
}
