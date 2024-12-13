package handler

import (
	"encoding/json"
	"go-project/internal/user/model"
	"go-project/internal/user/service"
	"net/http"
)

type AppointmentHandler struct {
	Service *service.AppointmentService
}

func NewAppointmentHandler(service *service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{Service: service}
}

func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointmentRequest model.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointmentRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Set status default untuk appointment yang diajukan oleh user
	appointmentRequest.Status = "pending"

	// Convert AppointmentRequest ke model Appointment yang lengkap
	appointment := model.Appointment{
		Name:          appointmentRequest.Name,
		PhoneNumber:   appointmentRequest.PhoneNumber,
		Email:         appointmentRequest.Email,
		DateOfBooking: appointmentRequest.DateOfBooking,
		Time:          appointmentRequest.Time,
		Status:        appointmentRequest.Status,
		PDFFile:       appointmentRequest.PDFFile,
		Img:           appointmentRequest.Img,
	}

	// Panggil service untuk membuat appointment
	if err := h.Service.CreateAppointment(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons sukses
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}
