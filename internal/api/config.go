package api

import (
	"github.com/kashifsoofi/movie-api/internal/config"
)

type ApiConfig struct {
	config.Database
	config.HTTPServer
}

func LoadConfig() (ApiConfig, error) {
	var cfg ApiConfig
	if err := config.Load(&cfg); err != nil {
		return ApiConfig{}, err
	}

	return cfg, nil
}
