package repository

import (
	"database/sql"
	"go-project/internal/admin/model"
)

// CommentRepository adalah interface yang mendefinisikan operasi-operasi terkait komentar.
type CommentRepository interface {
	// GetAllComments mengambil semua komentar dari database.
	GetAllComments() ([]model.Comment, error)

	// UpdateCommentStatus memperbarui status komentar berdasarkan ID komentar.
	UpdateCommentStatus(commentID int, status string) error

	// DeleteComment menghapus komentar berdasarkan ID komentar.
	DeleteComment(commentID int) error

	// CreateComment membuat komentar baru dan mengembalikannya.
	CreateComment(comment *model.Comment) (*model.Comment, error)
}

// commentRepository adalah implementasi konkret dari CommentRepository.
type commentRepository struct {
	db *sql.DB // Koneksi ke database
}

// NewCommentRepository adalah konstruktor untuk membuat instance baru dari commentRepository.
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}

// GetAllComments mengambil semua komentar dari database dan mengembalikannya dalam bentuk slice.
func (r *commentRepository) GetAllComments() ([]model.Comment, error) {
	// Query untuk mengambil semua komentar
	query := `SELECT id, article_id, username, email, comment, parent_id, status, created_at, updated_at FROM comments`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err // Mengembalikan error jika terjadi kesalahan saat query
	}
	defer rows.Close()

	var comments []model.Comment
	// Memproses setiap baris hasil query
	for rows.Next() {
		var comment model.Comment
		var parentID sql.NullInt32 // Menggunakan sql.NullInt32 untuk menangani nilai null
		if err := rows.Scan(&comment.ID, &comment.ArticleID, &comment.Username, &comment.Email,
			&comment.Comment, &parentID, &comment.Status, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err // Mengembalikan error jika terjadi kesalahan saat pemindaian data
		}
		// Jika parentID valid, set nilai parentID di comment
		if parentID.Valid {
			id := int(parentID.Int32)
			comment.ParentID = &id
		}
		comments = append(comments, comment) // Menambahkan komentar ke slice
	}

	return comments, nil // Mengembalikan slice komentar
}

// UpdateCommentStatus memperbarui status komentar berdasarkan ID komentar.
func (r *commentRepository) UpdateCommentStatus(commentID int, status string) error {
	query := `UPDATE comments SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, commentID)
	return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
}

// DeleteComment menghapus komentar berdasarkan ID komentar.
func (r *commentRepository) DeleteComment(commentID int) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.Exec(query, commentID)
	return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
}

// CreateComment menambahkan komentar baru ke database dan mengembalikan objek komentar yang baru dibuat.
func (r *commentRepository) CreateComment(comment *model.Comment) (*model.Comment, error) {
	query := `INSERT INTO comments (article_id, username, email, comment, parent_id, status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id, created_at, updated_at`
	// Menjalankan query dan mengembalikan ID, created_at, dan updated_at dari komentar yang baru dibuat
	err := r.db.QueryRow(query, comment.ArticleID, comment.Username, comment.Email, comment.Comment, comment.ParentID, comment.Status).
		Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err // Mengembalikan error jika terjadi kesalahan saat menambah data
	}

	return comment, nil // Mengembalikan komentar yang baru dibuat
}
