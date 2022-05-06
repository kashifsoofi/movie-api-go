package api

import "github.com/go-chi/chi/v5"

func (s *Server) routes() {
	s.router.Get("/health", s.handleGetHealth())

	s.router.Route("/movies", func(r chi.Router) {
		r.Get("/", s.handleListMovies())
		r.Post("/", s.handleCreateMovie())
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.handleGetMovie())
			r.Delete("/", s.handleDeleteMovie())
		})
	})
}
