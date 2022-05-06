package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

type HealthHandler struct {
	DB *sqlx.DB
}

type healthResponse struct {
	OK bool `json:"OK"`
}

func (hr healthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *HealthHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := healthResponse{OK: true}
		render.Render(w, r, health)
	}
}

func NewHealthHandler(db *sqlx.DB) *HealthHandler {
	return &HealthHandler{
		DB: db,
	}
}
