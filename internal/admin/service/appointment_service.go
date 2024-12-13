package service

import (
	"errors"
	"fmt"
	"go-project/internal/admin/model"
	"go-project/internal/admin/repository"
	"time"
)

// AppointmentService adalah layanan yang menyediakan logika bisnis untuk janji temu.
type AppointmentService struct {
	Repo *repository.AppointmentRepository // Repositori untuk operasi database terkait janji temu
}

// NewAppointmentService adalah konstruktor untuk membuat instance baru dari AppointmentService.
func NewAppointmentService(repo *repository.AppointmentRepository) *AppointmentService {
	return &AppointmentService{Repo: repo}
}

// ListStaff mengambil daftar staf yang tersedia dari repositori.
func (s *AppointmentService) ListStaff() ([]model.Staff, error) {
	// Mengambil daftar staf yang terdaftar di repositori
	return s.Repo.GetStaffList()
}

// CreateAppointment membuat janji temu baru dan menyimpannya ke dalam repositori.
func (s *AppointmentService) CreateAppointment(appointment *model.Appointment) error {
	// Memastikan format tanggal dan waktu valid
	parsedDateOfBooking, err := time.Parse(time.RFC3339, appointment.DateOfBooking.Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("invalid date_of_booking format: %v", err) // Mengembalikan error jika format tanggal tidak valid
	}
	appointment.DateOfBooking = parsedDateOfBooking

	parsedTime, err := time.Parse(time.RFC3339, appointment.Time.Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("invalid time format: %v", err) // Mengembalikan error jika format waktu tidak valid
	}
	appointment.Time = parsedTime

	// Validasi jika kolom wajib kosong
	if appointment.Name == "" || appointment.Email == "" || appointment.DateOfBooking.IsZero() || appointment.Time.IsZero() {
		return errors.New("all fields are required") // Mengembalikan error jika ada kolom wajib yang kosong
	}

	appointment.Status = "pending" // Menetapkan status janji temu menjadi "pending"
	// Menyimpan janji temu ke repositori
	return s.Repo.CreateAppointment(appointment)
}

// AssignHost menetapkan host (pemandu) untuk janji temu berdasarkan ID janji temu dan ID host.
func (s *AppointmentService) AssignHost(appointmentID, hostID int) error {
	// Memperbarui janji temu dengan host yang ditugaskan
	return s.Repo.UpdateAppointmentHost(appointmentID, hostID)
}

// UpdateStatus memperbarui status janji temu berdasarkan ID janji temu dan status yang baru.
func (s *AppointmentService) UpdateStatus(appointmentID int, status string) error {
	// Daftar status yang valid untuk janji temu
	validStatuses := map[string]bool{"confirmed": true, "pending": true, "cancelled": true}
	// Memeriksa apakah status yang diberikan valid
	if !validStatuses[status] {
		return errors.New("invalid status") // Mengembalikan error jika status tidak valid
	}
	// Memperbarui status janji temu di repositori
	return s.Repo.UpdateAppointmentStatus(appointmentID, status)
}
