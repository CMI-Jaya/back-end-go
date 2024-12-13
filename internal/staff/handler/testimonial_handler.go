package handler

import (
	"encoding/json"
	"go-project/internal/staff/model"
	"go-project/internal/staff/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TestimonialHandler struct {
	service service.TestimonialService
}

func NewTestimonialHandler(service service.TestimonialService) *TestimonialHandler {
	return &TestimonialHandler{service: service}
}

func (h *TestimonialHandler) CreateTestimonial(w http.ResponseWriter, r *http.Request) {
	var testimonial model.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&testimonial); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTestimonial(&testimonial); err != nil {
		http.Error(w, "Failed ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(testimonial)
}

func (h *TestimonialHandler) GetPendingTestimonials(w http.ResponseWriter, r *http.Request) {
	testimonials, err := h.service.GetPendingTestimonials()
	if err != nil {
		http.Error(w, "Failed lagi", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(testimonials)
}

func (h *TestimonialHandler) UpdateTestimonial(w http.ResponseWriter, r *http.Request) {
	var testimonial model.Testimonial
	if err := json.NewDecoder(r.Body).Decode(&testimonial); err != nil {
		http.Error(w, "Invalid Input !", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid testimonial ID", http.StatusBadRequest)
		return
	}
	testimonial.ID = id

	if err := h.service.UpdatePendingTestimonial(&testimonial); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(testimonial)
}

func (h *TestimonialHandler) DeleteTestimonial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid testimonial ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeletePendingTestimonial(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
