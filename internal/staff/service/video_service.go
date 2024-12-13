package service

import (
	"fmt"
	"go-project/internal/staff/model"
	"go-project/internal/staff/repository"
	"go-project/pkg/utils"
)

type VideoService struct {
	Repo             *repository.VideoRepository
	NotificationRepo *repository.NotificationRepository
}

// Konstruktor untuk VideoService
func NewVideoService(repo *repository.VideoRepository, notificationRepo *repository.NotificationRepository) *VideoService {
	if notificationRepo == nil {
		// Error handling jika notificationRepo tidak diinisialisasi
		panic("Notification repository is not initialized")
	}

	return &VideoService{
		Repo:             repo,
		NotificationRepo: notificationRepo, // Pastikan repository notifikasi diinisialisasi dengan benar
	}
}

func (s *VideoService) CreateVideo(video model.Video) error {
	// Validasi status hanya di sini
	validStatuses := map[string]bool{
		"pending approval": true,
	}

	// Jika status tidak valid, kembalikan error
	if !validStatuses[video.Status] && video.Status != "" {
		return fmt.Errorf("invalid status: %s", video.Status)
	}

	// Simpan video
	err := s.Repo.SaveVideo(video)
	if err != nil {
		return err
	}

	// Kirim notifikasi WhatsApp dan simpan ke database
	message := fmt.Sprintf("Video baru '%s' telah diunggah!", video.Title)
	err = utils.SendWhatsAppNotification("whatsapp:+6285707202183", message)
	if err != nil {
		return fmt.Errorf("failed to send WhatsApp notification: %w", err)
	}

	// Pastikan NotificationRepo tidak nil sebelum memanggil SaveNotification
	if s.NotificationRepo == nil {
		return fmt.Errorf("Notification repository is not initialized")
	}

	notification := model.Notification{
		UserID:  video.AuthorID,
		Type:    "video",
		Message: message,
		Status:  "unread",
	}
	return s.NotificationRepo.SaveNotification(notification)
}

func (s *VideoService) GetVideoByID(id int) (*model.Video, error) {
	return s.Repo.GetVideoByID(id)
}

func (s *VideoService) GetAllVideos() ([]model.Video, error) {
	return s.Repo.GetAllVideos()
}
