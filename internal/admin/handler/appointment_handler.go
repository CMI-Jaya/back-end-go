package handler

import (
	"encoding/json"
	"go-project/internal/admin/model"
	"go-project/internal/admin/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AppointmentHandler struct {
	Service *service.AppointmentService
}

// NewAppointmentHandler
// ---------------------
// Fungsi ini digunakan untuk menginisialisasi handler Appointment
// dengan menghubungkan ke layer service.
// Parameter:
// - service: Instance dari AppointmentService yang menyediakan logika bisnis.
// Return:
// - Pointer ke AppointmentHandler yang telah diinisialisasi.
func NewAppointmentHandler(service *service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{Service: service}
}

// GetStaffList
// ------------
// Fungsi ini digunakan untuk mendapatkan daftar staff dengan status "pending".
// Fungsi akan memanggil layer service untuk mengambil data.
// Jika berhasil, data akan dikembalikan dalam format JSON.
// Jika terjadi error, response HTTP akan mengembalikan status 500.
func (h *AppointmentHandler) GetStaffList(w http.ResponseWriter, r *http.Request) {
	staffList, err := h.Service.ListStaff()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(staffList)
}

// CreateAppointment
// ------------------
// Fungsi ini digunakan untuk membuat data appointment baru
// Parameter akan diterima dalam format JSON pada body request.
// Setelah data divalidasi, akan diteruskan ke layer service untuk diproses.

func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment model.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded Appointment: %+v\n", appointment)

	if appointment.Name == "" || appointment.Email == "" || appointment.DateOfBooking.IsZero() || appointment.Time.IsZero() {
		log.Println("One or more required fields are missing")
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateAppointment(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

// AssignHost
// -----------
// Fungsi ini digunakan untuk menetapkan host ke sebuah appointment.
//
// Parameter:
// - id (path variable): ID appointment yang akan diatur host-nya.
// - host_id (JSON body): ID host yang akan ditetapkan.
func (h *AppointmentHandler) AssignHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	var requestBody struct {
		HostID int `json:"host_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestBody.HostID <= 0 {
		http.Error(w, "Invalid host ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.AssignHost(appointmentID, requestBody.HostID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Host assigned successfully"))
}

// UpdateStatus
// -------------
// Fungsi ini digunakan untuk mengupdate status dari sebuah appointment.
//
// Parameter:
// - id (path variable): ID appointment yang statusnya akan diubah.
// - status (JSON body): Status baru yang akan diterapkan (approve/pending/rejected).

func (h *AppointmentHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}

	var data struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateStatus(appointmentID, data.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status updated successfully"))
}
