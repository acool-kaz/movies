package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Movies)
}
