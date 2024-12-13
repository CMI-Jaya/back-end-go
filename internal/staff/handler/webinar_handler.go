package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-project/internal/staff/service"
)

type WebinarHandler struct {
	service service.WebinarService
}

func NewWebinarHandler(service service.WebinarService) *WebinarHandler {
	return &WebinarHandler{service: service}
}

// GetAllWebinars handles GET requests to fetch all webinars
func (h *WebinarHandler) GetAllWebinars(w http.ResponseWriter, r *http.Request) {
	webinars, err := h.service.GetAllWebinars()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(webinars)
}

// GetWebinarByID handles GET requests to fetch a webinar by ID
func (h *WebinarHandler) GetWebinarByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing webinar ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid webinar ID", http.StatusBadRequest)
		return
	}

	webinar, err := h.service.GetWebinarByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(webinar)
}
