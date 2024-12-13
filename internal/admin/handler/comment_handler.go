package handler

import (
	"encoding/json"
	"go-project/internal/admin/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	service service.CommentService
}

// NewCommentHandler
// ------------------
// Fungsi ini digunakan untuk menginisialisasi handler Comment
// dengan menghubungkan ke layer service.
//
// Parameter:
// - service: Instance dari CommentService yang menyediakan logika bisnis.
//
// Return:
// - Pointer ke CommentHandler yang telah diinisialisasi.
func NewCommentHandler(service service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// GetAllComments
// ---------------
// Fungsi ini digunakan untuk mengambil semua komentar yang tersedia.

func (h *CommentHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.service.GetAllComments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// ApproveComment
// ---------------
// Fungsi ini digunakan untuk menyetujui komentar berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID komentar yang akan disetujui.

func (h *CommentHandler) ApproveComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	if err := h.service.ApproveComment(commentID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RejectComment
// --------------
// Fungsi ini digunakan untuk menolak komentar berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID komentar yang akan ditolak.

func (h *CommentHandler) RejectComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	if err := h.service.RejectComment(commentID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteComment
// --------------
// Fungsi ini digunakan untuk menghapus komentar berdasarkan ID.
//
// Parameter:
// - id (path parameter): ID komentar yang akan dihapus.

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	// Only the admin can delete comments, implement authorization check here
	if err := h.service.DeleteComment(commentID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateComment
// --------------
// Fungsi ini digunakan untuk membuat komentar baru.
//
// Parameter:
// - JSON body: Mengandung informasi komentar yang akan dibuat.

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment service.NewCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdComment, err := h.service.CreateComment(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdComment)
}

// ReplyComment
// -------------
// Fungsi ini digunakan untuk membalas komentar yang ada.
//
// Parameter:
// - id (path parameter): ID komentar yang akan dibalas.
// - JSON body: Mengandung informasi balasan komentar.

func (h *CommentHandler) ReplyComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var reply service.NewCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&reply); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set ParentID to the comment ID for the reply
	reply.ParentID = &commentID
	createdReply, err := h.service.CreateComment(reply)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdReply)
}
