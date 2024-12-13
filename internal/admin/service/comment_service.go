package service

import (
	"go-project/internal/admin/model"
	"go-project/internal/admin/repository"
)

// CommentService menyediakan layanan terkait komentar
type CommentService interface {
	GetAllComments() ([]model.Comment, error)                        // Mengambil semua komentar
	ApproveComment(commentID int) error                              // Menyetujui komentar
	RejectComment(commentID int) error                               // Menolak komentar
	DeleteComment(commentID int) error                               // Menghapus komentar
	CreateComment(comment NewCommentRequest) (*model.Comment, error) // Membuat komentar baru
}

type commentService struct {
	repo repository.CommentRepository // Repositori yang digunakan untuk operasi database komentar
}

// NewCommentService membuat instance baru dari CommentService
func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

// NewCommentRequest adalah struktur yang digunakan untuk permintaan pembuatan komentar baru
type NewCommentRequest struct {
	ArticleID int    `json:"article_id"`          // ID artikel terkait
	Username  string `json:"username"`            // Nama pengguna
	Email     string `json:"email"`               // Email pengguna
	Comment   string `json:"comment"`             // Isi komentar
	ParentID  *int   `json:"parent_id,omitempty"` // ID komentar induk, jika ada
}

// GetAllComments mengambil semua komentar
func (s *commentService) GetAllComments() ([]model.Comment, error) {
	return s.repo.GetAllComments() // Memanggil repositori untuk mengambil semua komentar
}

// ApproveComment menyetujui komentar
func (s *commentService) ApproveComment(commentID int) error {
	return s.repo.UpdateCommentStatus(commentID, "approved") // Memanggil repositori untuk memperbarui status komentar menjadi "approved"
}

// RejectComment menolak komentar
func (s *commentService) RejectComment(commentID int) error {
	return s.repo.UpdateCommentStatus(commentID, "rejected") // Memanggil repositori untuk memperbarui status komentar menjadi "rejected"
}

// DeleteComment menghapus komentar
func (s *commentService) DeleteComment(commentID int) error {
	return s.repo.DeleteComment(commentID) // Memanggil repositori untuk menghapus komentar berdasarkan ID
}

// CreateComment membuat komentar baru
func (s *commentService) CreateComment(req NewCommentRequest) (*model.Comment, error) {
	// Membuat objek komentar dari permintaan
	comment := model.Comment{
		ArticleID: req.ArticleID, // ID artikel yang terkait dengan komentar
		Username:  req.Username,  // Nama pengguna yang membuat komentar
		Email:     req.Email,     // Email pengguna yang membuat komentar
		Comment:   req.Comment,   // Isi komentar
		ParentID:  req.ParentID,  // ID komentar induk (jika ada)
		Status:    "pending",     // Status komentar awal adalah "pending"
	}

	// Memanggil repositori untuk menyimpan komentar
	return s.repo.CreateComment(&comment)
}
