package repository

import (
	"database/sql"
	"errors"
	"go-project/internal/staff/model"
)

type TestimonialRepository interface {
	CreateTestimonial(testimonial *model.Testimonial) error
	GetPendingTestimonials() ([]model.Testimonial, error)
	UpdatePendingTestimonial(testimonial *model.Testimonial) error
	DeletePendingTestimonial(id int) error
}

type testimonialRepository struct {
	db *sql.DB
}

func NewTestimonialRepository(db *sql.DB) TestimonialRepository {
	return &testimonialRepository{db: db}
}

func (r *testimonialRepository) CreateTestimonial(testimonial *model.Testimonial) error {
	query := `INSERT INTO testimonials (name, comment, photo_profile, category_id, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, 'pending', NOW(), NOW()) RETURNING id`
	err := r.db.QueryRow(query, testimonial.Name, testimonial.Comment, testimonial.PhotoProfile, testimonial.CategoryID).Scan(&testimonial.ID)
	return err
}

func (r *testimonialRepository) GetPendingTestimonials() ([]model.Testimonial, error) {
	query := `SELECT id, name, comment, photo_profile, category_id, status, created_at, updated_at
	FROM testimonials WHERE status = 'pending'`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonials []model.Testimonial
	for rows.Next() {
		var t model.Testimonial
		if err := rows.Scan(&t.ID, &t.Name, &t.Comment, &t.PhotoProfile, &t.CategoryID, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		testimonials = append(testimonials, t)
	}
	return testimonials, nil
}

func (r *testimonialRepository) UpdatePendingTestimonial(testimonial *model.Testimonial) error {
	query := `UPDATE testimonials SET name = $1, comment = $2, photo_profile = $3, category_id = $4, updated_at = NOW()
	WHERE id = $5 AND status = 'pending'`
	result, err := r.db.Exec(query, testimonial.Name, testimonial.Comment, testimonial.PhotoProfile, testimonial.CategoryID, testimonial.ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("No pending testimonial found for update !")
	}
	return nil
}

func (r *testimonialRepository) DeletePendingTestimonial(id int) error {
	query := `DELETE FROM testimonials WHERE id = $1 AND status = 'pending'`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("No pending testimonial !")
	}
	return nil
}
