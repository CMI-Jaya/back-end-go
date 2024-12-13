package repository

import (
	"database/sql"
	"go-project/internal/admin/model"
)

// AppointmentRepository adalah struct yang menyediakan metode untuk berinteraksi dengan tabel
// `appointments` di database. Struct ini memiliki properti DB yang menyimpan koneksi ke database.
type AppointmentRepository struct {
	DB *sql.DB
}

// NewAppointmentRepository adalah konstruktor yang membuat dan mengembalikan instance baru dari
// AppointmentRepository dengan koneksi database yang diberikan.
func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{DB: db}
}

// GetStaffList mengambil daftar staf dengan peran 'staff' dari tabel `users` dan mengembalikannya
// dalam bentuk slice dari model.Staff. Jika terjadi kesalahan saat query, akan mengembalikan error.
func (r *AppointmentRepository) GetStaffList() ([]model.Staff, error) {
	query := "SELECT id, name, email, role FROM users WHERE role = 'staff'"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var StaffList []model.Staff
	for rows.Next() {
		var staff model.Staff
		if err := rows.Scan(&staff.ID, &staff.Name, &staff.Email, &staff.Role); err != nil {
			return nil, err
		}
		StaffList = append(StaffList, staff)
	}
	return StaffList, nil
}

// CreateAppointment menyimpan data janji temu ke dalam tabel `appointments` dan mengembalikan
// error jika terjadi masalah saat penyimpanan. Fungsi ini juga mengembalikan ID janji temu yang baru
// yang disisipkan ke dalam tabel.
func (r *AppointmentRepository) CreateAppointment(appointment *model.Appointment) error {
	query := `INSERT INTO appointments (name, phone_number, email, date_of_booking, time, link_meet, host_id, status, pdf_file, img, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9 NOW(), NOW()) RETURNING id`
	err := r.DB.QueryRow(query, appointment.Name, appointment.PhoneNumber, appointment.Email, appointment.DateOfBooking,
		appointment.Time, appointment.LinkMeet, appointment.HostID, appointment.Status, appointment.PDFFile, appointment.Img).Scan(&appointment.ID)
	return err
}

// UpdateAppointmentHost memperbarui ID host pada janji temu tertentu berdasarkan appointmentID.
// Fungsi ini juga memperbarui waktu pembaruan (updated_at).
func (r *AppointmentRepository) UpdateAppointmentHost(appointmentID, hostID int) error {
	query := "UPDATE appointments SET host_id = $1, updated_at = NOW() WHERE id = $2"
	_, err := r.DB.Exec(query, hostID, appointmentID)
	return err
}

// UpdateAppointmentStatus memperbarui status janji temu berdasarkan appointmentID dan status yang
// diberikan. Fungsi ini juga memperbarui waktu pembaruan (updated_at).
func (r *AppointmentRepository) UpdateAppointmentStatus(appointmentID int, status string) error {
	query := "UPDATE appointments SET status = $1, updated_at = NOW() WHERE id = $2"
	_, err := r.DB.Exec(query, status, appointmentID)
	return err
}
