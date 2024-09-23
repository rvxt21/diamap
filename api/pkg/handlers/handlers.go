package handlers

import (
	"diamap/api/pkg/enteties"
	"encoding/json"
	"net/http"
)

type service interface {
	GetAllPodcasts() ([]enteties.PodcastShortInfo, error)
}

type Handler struct {
	Service service
}

func (h *Handler) GetAllPodcasts(w http.ResponseWriter, r *http.Request) {
	podcasts, err := h.Service.GetAllPodcasts()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(podcasts)
}
