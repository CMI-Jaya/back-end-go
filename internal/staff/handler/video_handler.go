package handler

import (
	"encoding/json"
	"go-project/internal/staff/model"
	"go-project/internal/staff/service"
	"net/http"
	"strconv"
)

// VideoHandler handles HTTP requests for videos
type VideoHandler struct {
	Service *service.VideoService
}

// UploadVideo handles the video upload request
func (h *VideoHandler) UploadVideo(w http.ResponseWriter, r *http.Request) {
	var video model.Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Buat video melalui service
	if err := h.Service.CreateVideo(video); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Video uploaded successfully and notification sent!"))
}

// GetVideoByID retrieves a video by ID
func (h *VideoHandler) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	// Extract the video ID from the URL query
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Call the service to get the video by ID
	video, err := h.Service.GetVideoByID(id)
	if err != nil {
		http.Error(w, "Video not found", http.StatusNotFound)
		return
	}

	// Respond with the video data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}

// GetAllVideos retrieves all videos
func (h *VideoHandler) GetAllVideos(w http.ResponseWriter, r *http.Request) {
	// Call the service to get all videos
	videos, err := h.Service.GetAllVideos()
	if err != nil {
		http.Error(w, "Error fetching videos", http.StatusInternalServerError)
		return
	}

	// Respond with the list of videos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}
