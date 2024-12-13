package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-project/internal/admin/model"
	"go-project/internal/admin/service"
)

type WebinarHandler struct {
	service service.WebinarService
}

func NewWebinarHandler(service service.WebinarService) *WebinarHandler {
	return &WebinarHandler{service: service}
}

func (h *WebinarHandler) CreateWebinar(w http.ResponseWriter, r *http.Request) {
	var webinar model.Webinar
	if err := json.NewDecoder(r.Body).Decode(&webinar); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateWebinar(&webinar); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Webinar created successfully", "id": strconv.Itoa(webinar.ID)})
}
