package main

import (
	"context"
	"os"

	"github.com/kashifsoofi/movie-api/internal/api"
)

func newServer() (*api.Server, error) {
	apiConfig, err := api.LoadConfig()
	if err != nil {
		return nil, err
	}

	server := api.NewServer(apiConfig.HTTPServer)
	server.GetRouter()
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
