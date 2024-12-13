package service

import (
	"go-project/internal/admin/model"
	"go-project/internal/admin/repository"
)

type TestimonialService interface {
	CreateTestimonial(testimonial *model.Testimonial) error        // Membuat testimonial baru
	GetAllTestimonials(status string) ([]model.Testimonial, error) // Mengambil semua testimonial berdasarkan status
	GetTestimonialByID(id int) (*model.Testimonial, error)         // Mengambil testimonial berdasarkan ID
	UpdateTestimonial(testimonial *model.Testimonial) error        // Memperbarui testimonial
	DeleteTestimonial(id int) error                                // Menghapus testimonial berdasarkan ID
	ApproveTestimonial(id int) error                               // Menyetujui testimonial
	RejectedTestimonial(id int) error                              // Menolak testimonial
}

type testimonialService struct {
	repo repository.TestimonialRepository // Repositori untuk operasi terkait testimonial
}

// NewTestimonialService membuat instance baru dari TestimonialService
func NewTestimonialService(repo repository.TestimonialRepository) TestimonialService {
	return &testimonialService{repo: repo}
}

// CreateTestimonial membuat testimonial baru dengan memanggil repositori
func (s *testimonialService) CreateTestimonial(testimonial *model.Testimonial) error {
	return s.repo.CreateTestimonial(testimonial)
}

// GetAllTestimonials mengambil semua testimonial berdasarkan status
func (s *testimonialService) GetAllTestimonials(status string) ([]model.Testimonial, error) {
	return s.repo.GetAllTestimonials(status)
}

// GetTestimonialByID mengambil testimonial berdasarkan ID
func (s *testimonialService) GetTestimonialByID(id int) (*model.Testimonial, error) {
	return s.repo.GetTestimonialByID(id)
}

// UpdateTestimonial memperbarui testimonial berdasarkan data yang diberikan
func (s *testimonialService) UpdateTestimonial(testimonial *model.Testimonial) error {
	return s.repo.UpdateTestimonial(testimonial)
}

// DeleteTestimonial menghapus testimonial berdasarkan ID
func (s *testimonialService) DeleteTestimonial(id int) error {
	return s.repo.DeleteTestimonial(id)
}

// ApproveTestimonial menyetujui testimonial dengan mengubah status menjadi "approved"
func (s *testimonialService) ApproveTestimonial(id int) error {
	return s.repo.UpdateStatus(id, "approved")
}

// RejectedTestimonial menolak testimonial dengan mengubah status menjadi "rejected"
func (s *testimonialService) RejectedTestimonial(id int) error {
	return s.repo.UpdateStatus(id, "rejected")
}
