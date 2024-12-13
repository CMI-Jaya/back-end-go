package repository

import (
	"database/sql"
	"go-project/internal/user/model"
)

type AppointmentRepository struct {
	DB *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{DB: db}
}

func (r *AppointmentRepository) CreateAppointment(appointment *model.Appointment) error {
	query := `INSERT INTO appointments 
        (name, phone_number, email, date_of_booking, time, status, pdf_file, img, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()) RETURNING id`

	// Tidak memasukkan link_meet karena itu hanya diatur oleh admin
	return r.DB.QueryRow(query,
		appointment.Name,
		appointment.PhoneNumber,
		appointment.Email,
		appointment.DateOfBooking,
		appointment.Time,
		appointment.Status,
		appointment.PDFFile,
		appointment.Img).Scan(&appointment.ID)
}

func (r *AppointmentRepository) GetAppointments() ([]model.Appointment, error) {
	query := `SELECT id, name, email, date_of_booking, time, pdf_file, img, status,
              FROM appointments WHERE status = 'pending'`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []model.Appointment
	for rows.Next() {
		var appointment model.Appointment
		if err := rows.Scan(
			&appointment.ID,
			&appointment.Name,
			&appointment.Email,
			&appointment.DateOfBooking,
			&appointment.Time,
			&appointment.PDFFile,
			&appointment.Img,
			&appointment.Status,
		); err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}
