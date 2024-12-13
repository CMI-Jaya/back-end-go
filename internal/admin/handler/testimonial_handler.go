package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-project/internal/admin/model"
	"go-project/internal/admin/service"

	"github.com/gorilla/mux"
)

type TestimonialHandler struct {
	service service.TestimonialService
}

// NewTestimonialHandler
// ----------------------
// Fungsi ini digunakan untuk menginisialisasi handler Testimonial
// dengan menghubungkan ke layer service.
//
// Parameter:
// - service: Instance dari TestimonialService yang menyediakan logika bisnis.
//
// Return:
// - Pointer ke TestimonialHandler yang telah diinisialisasi.
func NewTestimonialHandler(service service.TestimonialService) *TestimonialHandler {
	return &TestimonialHandler{service: service}
}

// CreateTestimonial
// ------------------
// Fungsi ini digunakan untuk membuat testimonial baru.
//
// Parameter:
// - JSON body: Mengandung informasi testimonial yang akan dibuat.

func (h *TestimonialHandler) CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	var testimonial model.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&testimonial); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTestimonial(&testimonial); err != nil {
		http.Error(w, "Failed to create testimonial", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(testimonial)
}

// GetAllTestimonials
// -------------------
// Fungsi ini digunakan untuk mengambil semua testimonial berdasarkan status.
//
// Query Parameter:
// - status (opsional): Filter berdasarkan status testimonial.

func (h *TestimonialHandler) GetAllTestimonials(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	testimonials, err := h.service.GetAllTestimonials(status)
	if err != nil {
		http.Error(w, "Failed to fetch testimonials", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(testimonials)
}

// GetTestimonialByID
// -------------------
// Fungsi ini digunakan untuk mengambil detail testimonial berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID testimonial yang akan diambil.

func (h *TestimonialHandler) GetTestimonialByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid testimonial ID", http.StatusBadRequest)
		return
	}

	testimonial, err := h.service.GetTestimonialByID(id)
	if err != nil {
		http.Error(w, "Testimonial not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(testimonial)
}

// UpdateTestimonial
// ------------------
// Fungsi ini digunakan untuk memperbarui testimonial berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID testimonial yang akan diperbarui.
// - JSON body: Mengandung informasi testimonial yang baru.

func (h *TestimonialHandler) UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
	var testimonial model.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&testimonial); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid testimonial ID", http.StatusBadRequest)
		return
	}
	testimonial.ID = id

	if err := h.service.UpdateTestimonial(&testimonial); err != nil {
		http.Error(w, "Failed to update testimonial", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(testimonial)
}

// DeleteTestimonial
// ------------------
// Fungsi ini digunakan untuk menghapus testimonial berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID testimonial yang akan dihapus.

func (h *TestimonialHandler) DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid testimonial ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTestimonial(id); err != nil {
		http.Error(w, "Failed to delete testimonial", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ApproveTestimonial
// -------------------
// Fungsi ini digunakan untuk menyetujui testimonial berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID testimonial yang akan disetujui.

func (h *TestimonialHandler) ApproveTestimonial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid testimonial ID", http.StatusBadRequest)
		return
	}

	if err := h.service.ApproveTestimonial(id); err != nil {
		http.Error(w, "Failed to approve testimonial", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RejectTestimonial
// ------------------
// Fungsi ini digunakan untuk menolak testimonial berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID testimonial yang akan ditolak.

func (h *TestimonialHandler) RejectTestimonial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid testimonial ID", http.StatusBadRequest)
		return
	}

	if err := h.service.RejectedTestimonial(id); err != nil {
		http.Error(w, "Failed to reject testimonial", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
