package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/acool-kaz/movies/models"
)

func (h *Handler) createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	if err := movie.Validation(); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	movie.ID = strconv.Itoa(len(h.Movies) + 1)
	h.Movies = append(h.Movies, movie)
	json.NewEncoder(w).Encode(h.Movies)
}
