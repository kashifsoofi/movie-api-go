package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/kashifsoofi/movie-api/internal/api/handler"
)

func (s *Server) GetRouter() *chi.Mux {
	r := chi.NewRouter()

	hh := handler.NewHealthHandler(nil)
	r.Get("/health", hh.Get())

	s.router = r

	return s.router
}
