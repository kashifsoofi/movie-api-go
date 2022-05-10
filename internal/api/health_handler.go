package api

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
)

type healthResponse struct {
	OK      bool `json:"ok"`
	StoreOK bool `json:"store_ok"`
}

func (hr healthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) handleGetHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := healthResponse{OK: true}
		health.StoreOK = s.storeOK(r.Context())

		render.Render(w, r, health)
	}
}

func (s *Server) storeOK(ctx context.Context) bool {
	err := s.store.Connect(ctx)
	if err != nil {
		return false
	}
	defer s.store.Close()

	return true
}
