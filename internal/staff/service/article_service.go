package service

import (
	"fmt"
	"go-project/internal/staff/model"
	"go-project/internal/staff/repository"
	"go-project/pkg/utils"
)

type ArticleService struct {
	Repo             *repository.ArticleRepository
	NotificationRepo *repository.NotificationRepository // Tambahkan untuk menyimpan notifikasi
}

func (s *ArticleService) CreateArticle(article model.Article) error {
	if s.NotificationRepo == nil {
		return fmt.Errorf("Notification repository is not initialized")
	}
	// Validasi status
	validStatuses := map[string]bool{
		"pending approval": true,
	}

	if !validStatuses[article.Status] {
		return fmt.Errorf("invalid status: %s", article.Status)
	}

	// Simpan artikel
	err := s.Repo.SaveArticle(article)
	if err != nil {
		return err
	}

	// Kirim notifikasi WhatsApp dan simpan ke database
	message := fmt.Sprintf("Artikel baru '%s' telah diunggah!", article.Title)
	err = utils.SendWhatsAppNotification("whatsapp:+6285707202183", message)
	if err != nil {
		return fmt.Errorf("failed to send WhatsApp notification: %w", err)
	}

	notification := model.Notification{
		UserID:  article.AuthorID,
		Type:    "article",
		Message: message,
		Status:  "unread",
	}
	return s.NotificationRepo.SaveNotification(notification)
}

func (s *ArticleService) GetArticleByID(id int) (*model.Article, error) {
	return s.Repo.GetArticleByID(id) // Mendapatkan artikel berdasarkan ID
}

func (s *ArticleService) GetAllArticles() ([]model.Article, error) {
	return s.Repo.GetAllArticles() // Mendapatkan semua artikel
}
