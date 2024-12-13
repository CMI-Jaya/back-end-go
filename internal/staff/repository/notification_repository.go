package repository

import (
	"database/sql"
	"fmt"
	"go-project/internal/staff/model"
)

type NotificationRepository struct {
	DB *sql.DB
}

// Fungsi untuk menyimpan notifikasi
func (r *NotificationRepository) SaveNotification(notification model.Notification) error {
	query := `INSERT INTO notifications (user_id, type, message, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := r.DB.Exec(query, notification.UserID, notification.Type, notification.Message, notification.Status)
	return err
}

// Fungsi untuk mendapatkan user berdasarkan userID
func (r *NotificationRepository) GetUserByID(userID int) (model.User, error) {
	var user model.User
	// Menjalankan query untuk mengambil data user berdasarkan userID
	query := `SELECT id, role, name, email FROM users WHERE id = $1`
	err := r.DB.QueryRow(query, userID).Scan(&user.ID, &user.Role, &user.Name, &user.Email)
	if err != nil {
		return user, fmt.Errorf("failed to get user with ID %d: %w", userID, err)
	}
	return user, nil
}
