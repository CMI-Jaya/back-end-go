package service

import (
	"go-project/internal/user/model"
	"go-project/internal/user/repository"
)

type AppointmentService struct {
	Repo *repository.AppointmentRepository
}

func NewAppointmentService(repo *repository.AppointmentRepository) *AppointmentService {
	return &AppointmentService{Repo: repo}
}

func (s *AppointmentService) CreateAppointment(appointment *model.Appointment) error {
	// Set status appointment ke "pending"
	appointment.Status = "pending"
	return s.Repo.CreateAppointment(appointment)
}

func (s *AppointmentService) ListAppointments() ([]model.Appointment, error) {
	return s.Repo.GetAppointments()
}
