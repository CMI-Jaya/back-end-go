package service

import (
	"go-project/internal/staff/model"
	"go-project/internal/staff/repository"
)

type TestimonialService interface {
	CreateTestimonial(testimonial *model.Testimonial) error
	GetPendingTestimonials() ([]model.Testimonial, error)
	UpdatePendingTestimonial(testimonial *model.Testimonial) error
	DeletePendingTestimonial(id int) error
}

type testimonialService struct {
	repo repository.TestimonialRepository
}

func NewTestimonialService(repo repository.TestimonialRepository) TestimonialService {
	return &testimonialService{repo: repo}
}

func (s *testimonialService) CreateTestimonial(testimonial *model.Testimonial) error {
	return s.repo.CreateTestimonial(testimonial)
}

func (s *testimonialService) GetPendingTestimonials() ([]model.Testimonial, error) {
	return s.repo.GetPendingTestimonials()
}

func (s *testimonialService) UpdatePendingTestimonial(testimonial *model.Testimonial) error {
	return s.repo.UpdatePendingTestimonial(testimonial)
}

func (s *testimonialService) DeletePendingTestimonial(id int) error {
	return s.repo.DeletePendingTestimonial(id)
}
