package handlers

import "github.com/gorilla/mux"

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/podcasts", h.GetAllPodcasts).Methods("GET")
}
