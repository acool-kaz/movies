package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range h.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}
