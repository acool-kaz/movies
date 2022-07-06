package handler

import (
	"github.com/acool-kaz/movies/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Movies []models.Movie
}

func NewHandler() *Handler {
	return &Handler{
		Router: mux.NewRouter(),
	}
}

func (h *Handler) InitEndpoits() {
	h.Router.HandleFunc("/movies", h.getMovies).Methods("GET")
	h.Router.HandleFunc("/movies/{id}", h.getMovie).Methods("GET")
	h.Router.HandleFunc("/movies", h.createMovie).Methods("POST")
	h.Router.HandleFunc("/movies/{id}", h.updateMovie).Methods("PUT")
	h.Router.HandleFunc("/movies/{id}", h.deleteMovie).Methods("DELETE")
}
