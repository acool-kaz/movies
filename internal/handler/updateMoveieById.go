package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/acool-kaz/movies/models"
	"github.com/gorilla/mux"
)

func (h *Handler) updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range h.Movies {
		if item.ID == params["id"] {
			var movie models.Movie
			json.NewDecoder(r.Body).Decode(&movie)
			if err := movie.Validation(); err != nil {
				fmt.Fprint(w, err.Error())
				return
			}
			movie.ID = params["id"]
			h.Movies[index] = movie
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}
