package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range h.Movies {
		if item.ID == params["id"] {
			h.Movies = append(h.Movies[:index], h.Movies[index+1:]...)
			for i := range h.Movies[index:] {
				id, _ := strconv.Atoi(h.Movies[index:][i].ID)
				id--
				h.Movies[index:][i].ID = strconv.Itoa(id)
			}
			break
		}
	}
	json.NewEncoder(w).Encode(h.Movies)
}
