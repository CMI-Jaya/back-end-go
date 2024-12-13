package service

import (
	"fmt"
	"go-project/internal/staff/model"
	"go-project/internal/staff/repository"
	"go-project/pkg/utils"
	"os"
)

type NotificationService struct {
	Repo *repository.NotificationRepository
}

// Fungsi untuk mengecek apakah user adalah admin berdasarkan userID
func (s *NotificationService) IsAdmin(userID int) (bool, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return false, fmt.Errorf("failed to get user: %w", err)
	}
	// Memeriksa apakah role pengguna adalah "admin"
	return user.Role == "admin", nil
}

func (s *NotificationService) SaveAndSendNotification(userID int, notificationType, message string) error {
	// Pastikan repository tidak nil
	if s.Repo == nil {
		return fmt.Errorf("notification repository is not initialized")
	}

	// Simpan notifikasi ke database
	notification := model.Notification{
		UserID:  userID,
		Type:    notificationType,
		Message: message,
		Status:  "unread",
	}
	err := s.Repo.SaveNotification(notification)
	if err != nil {
		return fmt.Errorf("failed to save notification: %w", err)
	}

	// Cek apakah user adalah admin sebelum mengirim notifikasi
	isAdmin, err := s.IsAdmin(userID)
	if err != nil {
		return fmt.Errorf("failed to check if user is admin: %w", err)
	}

	if isAdmin {
		to := os.Getenv("ADMIN_WHATSAPP_NUMBER")
		if to == "" {
			return fmt.Errorf("admin WhatsApp number is not configured")
		}

		err = utils.SendWhatsAppNotification(to, message)
		if err != nil {
			return fmt.Errorf("failed to send WhatsApp notification: %w", err)
		}
	}

	return nil
}
