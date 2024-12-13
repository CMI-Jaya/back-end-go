package service

import (
	"go-project/internal/staff/model"
	"go-project/internal/staff/repository"
)

type AppointmentService struct {
	Repo *repository.AppointmentRepository
}

func NewAppointmentService(repo *repository.AppointmentRepository) *AppointmentService {
	return &AppointmentService{Repo: repo}
}

func (s *AppointmentService) CreateAppointment(appointment *model.Appointment) error {
	appointment.Status = "pending"
	return s.Repo.CreateAppointment(appointment)
}

func (s *AppointmentService) ListAppointments() ([]model.Appointment, error) {
	return s.Repo.GetAppointments()
}
