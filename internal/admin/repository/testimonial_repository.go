package repository

import (
	"database/sql"
	"errors"
	"go-project/internal/admin/model"
)

// TestimonialRepository adalah interface yang mendefinisikan operasi-operasi terkait testimonial.
type TestimonialRepository interface {
	// CreateTestimonial membuat testimonial baru dan menyimpannya dalam database.
	CreateTestimonial(testimonial *model.Testimonial) error

	// GetAllTestimonials mengambil semua testimonial dengan status tertentu.
	GetAllTestimonials(status string) ([]model.Testimonial, error)

	// GetTestimonialByID mengambil testimonial berdasarkan ID.
	GetTestimonialByID(id int) (*model.Testimonial, error)

	// UpdateTestimonial memperbarui testimonial berdasarkan objek testimonial yang diberikan.
	UpdateTestimonial(testimonial *model.Testimonial) error

	// DeleteTestimonial menghapus testimonial berdasarkan ID.
	DeleteTestimonial(id int) error

	// UpdateStatus memperbarui status testimonial berdasarkan ID.
	UpdateStatus(id int, status string) error
}

// testimonialRepository adalah implementasi konkret dari TestimonialRepository.
type testimonialRepository struct {
	db *sql.DB // Koneksi ke database
}

// NewTestimonialRepository adalah konstruktor untuk membuat instance baru dari testimonialRepository.
func NewTestimonialRepository(db *sql.DB) TestimonialRepository {
	return &testimonialRepository{db: db}
}

// CreateTestimonial membuat testimonial baru dan menyimpannya ke dalam database.
// Status testimonial awalnya adalah "pending".
func (r *testimonialRepository) CreateTestimonial(testimonial *model.Testimonial) error {
	query := `INSERT INTO testimonials (name, comment, photo_profile, category_id, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`
	err := r.db.QueryRow(query, testimonial.Name, testimonial.Comment, testimonial.PhotoProfile, testimonial.CategoryID, "pending").Scan(&testimonial.ID)
	return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
}

// GetAllTestimonials mengambil semua testimonial berdasarkan status tertentu (misalnya "approved", "pending").
func (r *testimonialRepository) GetAllTestimonials(status string) ([]model.Testimonial, error) {
	var testimonials []model.Testimonial
	query := `SELECT id, name, comment, photo_profile, category_id, status, created_at, updated_at FROM testimonials WHERE status = $1`
	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, err // Mengembalikan error jika terjadi kesalahan saat query
	}
	defer rows.Close()

	// Memproses setiap baris hasil query
	for rows.Next() {
		var t model.Testimonial
		if err := rows.Scan(&t.ID, &t.Name, &t.Comment, &t.PhotoProfile, &t.CategoryID, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err // Mengembalikan error jika terjadi kesalahan saat pemindaian data
		}
		testimonials = append(testimonials, t) // Menambahkan testimonial ke slice
	}
	return testimonials, nil // Mengembalikan slice testimonial
}

// GetTestimonialByID mengambil testimonial berdasarkan ID.
func (r *testimonialRepository) GetTestimonialByID(id int) (*model.Testimonial, error) {
	query := `SELECT id, name, comment, photo_profile, category_id, status, created_at, updated_at FROM testimonials WHERE id = $1`
	var t model.Testimonial
	err := r.db.QueryRow(query, id).Scan(&t.ID, &t.Name, &t.Comment, &t.PhotoProfile, &t.CategoryID, &t.Status, &t.CreatedAt, &t.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("Testimonials Not Found !") // Mengembalikan error jika testimonial tidak ditemukan
	}
	return &t, err // Mengembalikan testimonial dan error (jika ada)
}

// UpdateTestimonial memperbarui testimonial berdasarkan objek testimonial yang diberikan.
func (r *testimonialRepository) UpdateTestimonial(testimonial *model.Testimonial) error {
	query := `UPDATE testimonials SET name = $1, comment = $2, photo_profile = $3, category_id = $4, updated_at = NOW() WHERE id = $5`
	_, err := r.db.Exec(query, testimonial.Name, testimonial.Comment, testimonial.PhotoProfile, testimonial.CategoryID, testimonial.ID)
	return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
}

// DeleteTestimonial menghapus testimonial berdasarkan ID.
func (r *testimonialRepository) DeleteTestimonial(id int) error {
	query := `DELETE FROM testimonials WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
}

// UpdateStatus memperbarui status testimonial berdasarkan ID.
func (r *testimonialRepository) UpdateStatus(id int, status string) error {
	query := `UPDATE testimonials SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err // Mengembalikan error jika terjadi kesalahan saat eksekusi query
}
