package handler

import (
	"encoding/json"
	"go-project/internal/admin/repository"
	"go-project/internal/admin/service"
	"log"
	"net/http"
	"strconv"
)

// VideoHandler adalah struct yang menyediakan metode untuk menangani permintaan HTTP terkait video.
// Struct ini berisi referensi ke VideoService untuk berinteraksi dengan logika bisnis.
type VideoHandler struct {
	Service *service.VideoService
}

// NewVideoHandler adalah konstruktor yang membuat dan mengembalikan instance baru dari VideoHandler
// dengan VideoService yang diberikan.
func NewVideoHandler(service *service.VideoService) *VideoHandler {
	return &VideoHandler{Service: service}
}

// CreateVideo menangani pembuatan video baru. Fungsi ini membaca payload JSON dari body permintaan,
// memvalidasi field yang dibutuhkan, dan memanggil service untuk menyimpan video ke dalam database.
// Jika berhasil, ID video yang baru dibuat akan dikembalikan dalam respon.
func (h *VideoHandler) CreateVideo(w http.ResponseWriter, r *http.Request) {
	var video repository.Video

	// Decode payload JSON ke dalam struct video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		http.Error(w, "Payload JSON tidak valid", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded video: %+v\n", video)

	// Validasi field yang dibutuhkan
	if video.Title == "" || video.Description == "" || video.LinkVideo == "" || video.CategoryID == 0 {
		http.Error(w, "Semua field (title, description, link_video, category_id) wajib diisi", http.StatusBadRequest)
		return
	}

	// Simpan video ke dalam database
	id, err := h.Service.CreateVideo(video)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respon dengan status berhasil
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// GetAllVideos menangani pengambilan semua video. Fungsi ini memanggil service untuk mendapatkan daftar
// video dari database dan mengembalikan hasilnya dalam bentuk respon JSON.
func (h *VideoHandler) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := h.Service.GetAllVideos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}

// GetVideoByID menangani pengambilan video berdasarkan ID. Fungsi ini mengekstrak ID dari query
// parameter, memvalidasinya, dan memanggil service untuk mengambil detail video. Hasilnya
// dikembalikan dalam bentuk respon JSON atau pesan kesalahan jika video tidak ditemukan.
func (h *VideoHandler) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	video, err := h.Service.GetVideoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}

// UpdateVideo menangani pembaruan video yang sudah ada. Fungsi ini mengekstrak ID video dari query
// parameter, mendecode payload JSON dari body permintaan, dan memvalidasi field yang dibutuhkan.
// Video kemudian diperbarui di dalam database, dan video yang diperbarui dikembalikan dalam respon.
func (h *VideoHandler) UpdateVideo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var video repository.Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	video.ID = id
	// Validasi field yang dibutuhkan
	if video.Title == "" || video.Description == "" || video.LinkVideo == "" || video.CategoryID == 0 {
		http.Error(w, "Semua field (title, description, link_video, category_id) wajib diisi", http.StatusBadRequest)
		return
	}

	// Pembaruan video di dalam database
	if err := h.Service.UpdateVideo(video); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(video)
}

// DeleteVideo menangani penghapusan video. Fungsi ini mengekstrak ID video dari query
// parameter, memvalidasinya, dan memanggil service untuk menghapus video dari database.
// Jika berhasil, fungsi ini akan merespon dengan status "No Content".
func (h *VideoHandler) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	// Menghapus video dari database
	if err := h.Service.DeleteVideo(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
