package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/kashifsoofi/movie-api/internal/store"
)

type movieResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"release_date"`
	TicketPrice float64   `json:"ticket_price"`
}

func NewMovieResponse(m store.Movie) movieResponse {
	return movieResponse{
		ID:          m.ID,
		Title:       m.Title,
		Director:    m.Director,
		ReleaseDate: m.ReleaseDate,
		TicketPrice: m.TicketPrice,
	}
}

func (hr movieResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewMovieListResponse(movies []store.Movie) []render.Renderer {
	list := []render.Renderer{}
	for _, movie := range movies {
		mr := NewMovieResponse(movie)
		list = append(list, mr)
	}
	return list
}

func (s *Server) handleListMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, _ := s.store.Movies().GetAll()
		render.RenderList(w, r, NewMovieListResponse(movies))
	}
}

func (s *Server) handleGetMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			w.WriteHeader(400)
			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		movie, err := s.store.Movies().GetByID(id)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		mr := NewMovieResponse(movie)
		render.Render(w, r, mr)
	}
}

type movieRequest struct {
	Title       string    `json:"title"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"release_date"`
	TicketPrice float64   `json:"ticket_price"`
}

func (mr *movieRequest) Bind(r *http.Request) error {
	return nil
}

func (s *Server) handleCreateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &movieRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		movie := store.Movie{
			ID:          uuid.New(),
			Title:       data.Title,
			Director:    data.Director,
			ReleaseDate: data.ReleaseDate,
			TicketPrice: data.TicketPrice,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		}
		s.store.Movies().Create(movie)

		mr := NewMovieResponse(movie)
		render.Render(w, r, mr)
	}
}

func (s *Server) handleDeleteMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			w.WriteHeader(400)
			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		movie, err := s.store.Movies().Delete(id)
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		mr := NewMovieResponse(movie)
		render.Render(w, r, mr)
	}
}
