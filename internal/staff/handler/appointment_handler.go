package handler

import (
	"encoding/json"
	"go-project/internal/staff/model"
	"go-project/internal/staff/service"
	"log"
	"net/http"
)

type AppointmentHandler struct {
	Service *service.AppointmentService
}

func NewAppointmentHandler(service *service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{Service: service}
}

func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment model.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded Appointment: %+v\n", appointment)

	if err := h.Service.CreateAppointment(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

func (h *AppointmentHandler) ListAppointments(w http.ResponseWriter, r *http.Request) {
	appointments, err := h.Service.ListAppointments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}
