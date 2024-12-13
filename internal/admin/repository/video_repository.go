package repository

import (
	"database/sql"
	"errors"
	"time"
)

// Video adalah struktur yang mendefinisikan skema untuk entitas video.
type Video struct {
	ID              int    `json:"id,omitempty"`               // ID video
	Title           string `json:"title"`                      // Judul video
	Description     string `json:"description"`                // Deskripsi video
	LinkVideo       string `json:"link_video"`                 // Tautan video
	CategoryID      int    `json:"category_id"`                // ID kategori video
	AuthorID        int    `json:"author_id"`                  // ID penulis
	MetaTitle       string `json:"meta_title,omitempty"`       // Meta title untuk SEO
	MetaDescription string `json:"meta_description,omitempty"` // Meta description untuk SEO
	Status          string `json:"status"`                     // Status video (e.g., "approval", "pending approval", "rejected")
	CreatedAt       string `json:"-"`                          // Waktu pembuatan video
	UpdatedAt       string `json:"-"`                          // Waktu pembaruan video
}

// VideoRepository adalah struktur untuk repositori video yang berisi fungsi-fungsi untuk interaksi dengan database.
type VideoRepository struct {
	DB *sql.DB // Koneksi ke database
}

// NewVideoRepository adalah konstruktor untuk membuat instance baru dari VideoRepository.
func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{DB: db}
}

// validVideoStatus memeriksa apakah status yang diberikan valid untuk video.
func validVideoStatus(status string) bool {
	validStatuses := []string{"approval", "pending approval", "rejected"}
	for _, s := range validStatuses {
		if status == s {
			return true
		}
	}
	return false
}

// UpdateVideoStatus memperbarui status video berdasarkan ID.
func (repo *VideoRepository) UpdateVideoStatus(id int, status string) error {
	// Memeriksa apakah status yang diberikan valid
	if !validVideoStatus(status) {
		return errors.New("invalid status") // Mengembalikan error jika status tidak valid
	}

	query := `UPDATE videos SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := repo.DB.Exec(query, status, id) // Eksekusi query untuk memperbarui status
	return err
}

// Create membuat video baru dan menyimpannya ke dalam database.
func (repo *VideoRepository) Create(video Video) (int, error) {
	// Memeriksa apakah category_id valid
	if video.CategoryID == 0 {
		return 0, errors.New("category_id cannot be empty or zero")
	}

	// Query untuk memasukkan video baru ke database
	query := `INSERT INTO videos (title, description, link_video, category_id, meta_title, meta_description, status, author_id, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	var id int
	// Eksekusi query untuk menyimpan video dan mengembalikan ID-nya
	err := repo.DB.QueryRow(query, video.Title, video.Description, video.LinkVideo, video.CategoryID, video.MetaTitle, video.MetaDescription, video.Status, video.AuthorID,
		time.Now(), time.Now()).Scan(&id)

	return id, err // Mengembalikan ID video dan error jika ada
}

// GetAll mengambil semua video yang ada di database.
func (repo *VideoRepository) GetAll() ([]Video, error) {
	query := `SELECT id, title, description, link_video, category_id, meta_title, meta_description, created_at, updated_at FROM videos`
	rows, err := repo.DB.Query(query) // Eksekusi query untuk mengambil semua video
	if err != nil {
		return nil, err // Mengembalikan error jika terjadi kesalahan saat query
	}
	defer rows.Close()

	var videos []Video
	// Memproses setiap baris hasil query
	for rows.Next() {
		var video Video
		err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.LinkVideo, &video.CategoryID, &video.MetaTitle, &video.MetaDescription, &video.CreatedAt, &video.UpdatedAt)
		if err != nil {
			return nil, err // Mengembalikan error jika terjadi kesalahan saat pemindaian data
		}
		videos = append(videos, video) // Menambahkan video ke slice
	}
	return videos, nil // Mengembalikan slice video
}

// GetByID mengambil video berdasarkan ID.
func (repo *VideoRepository) GetByID(id int) (*Video, error) {
	query := `SELECT id, title, description, link_video, category_id, meta_title, meta_description, created_at, updated_at FROM videos WHERE id = $1`
	row := repo.DB.QueryRow(query, id) // Eksekusi query untuk mengambil video berdasarkan ID

	var video Video
	// Memindai hasil query ke dalam objek video
	err := row.Scan(&video.ID, &video.Title, &video.Description, &video.LinkVideo, &video.CategoryID, &video.MetaTitle, &video.MetaDescription, &video.CreatedAt, &video.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("Video Not Found!") // Mengembalikan error jika video tidak ditemukan
	}
	return &video, err // Mengembalikan video dan error (jika ada)
}

// Update memperbarui data video yang ada berdasarkan objek video yang diberikan.
func (repo *VideoRepository) Update(video Video) error {
	query := `UPDATE videos SET title = $1, description = $2, link_video = $3, category_id = $4, meta_title = $5, meta_description = $6, updated_at = $7 WHERE id = $8`
	_, err := repo.DB.Exec(query, video.Title, video.Description, video.LinkVideo, video.CategoryID, video.MetaTitle, video.MetaDescription, time.Now(), video.ID)
	return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
}

// Delete menghapus video berdasarkan ID.
func (repo *VideoRepository) Delete(id int) error {
	query := `DELETE FROM videos WHERE id = $1`
	_, err := repo.DB.Exec(query, id) // Eksekusi query untuk menghapus video berdasarkan ID
	return err
}
