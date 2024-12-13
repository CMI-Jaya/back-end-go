package service

import (
	"errors"
	"go-project/internal/admin/model"
	"go-project/internal/admin/repository"
	"net/url"
)

// ArticleService menyediakan logika bisnis terkait artikel
type ArticleService interface {
	CreateArticle(article *model.Article) error    // Fungsi untuk membuat artikel baru
	GetArticleByID(id int) (*model.Article, error) // Fungsi untuk mengambil artikel berdasarkan ID
	UpdateArticle(article *model.Article) error    // Fungsi untuk memperbarui artikel yang ada
	DeleteArticle(id int) error                    // Fungsi untuk menghapus artikel berdasarkan ID
	GetAllArticles() ([]model.Article, error)      // Fungsi untuk mengambil semua artikel
}

type articleService struct {
	repo repository.ArticleRepository // Repositori untuk operasi database terkait artikel
}

// NewArticleService membuat instance baru dari ArticleService
func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

// ValidateVideoURL memeriksa apakah URL yang diberikan valid
func (s *articleService) ValidateVideoURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr) // Memeriksa apakah URL valid
	return err == nil
}

// CreateArticle menyisipkan artikel baru ke dalam database
func (s *articleService) CreateArticle(article *model.Article) error {
	// Validasi URL video
	if !s.ValidateVideoURL(article.LinkVideo) {
		return errors.New("invalid video URL") // Mengembalikan error jika URL video tidak valid
	}

	// Memanggil lapisan repositori untuk membuat artikel
	return s.repo.CreateArticle(article)
}

// UpdateArticle memperbarui artikel yang ada
func (s *articleService) UpdateArticle(article *model.Article) error {
	// Validasi URL video jika URL video diberikan
	if article.LinkVideo != "" && !s.ValidateVideoURL(article.LinkVideo) {
		return errors.New("invalid video URL") // Mengembalikan error jika URL video tidak valid
	}

	// Memanggil repositori untuk memperbarui artikel
	return s.repo.UpdateArticle(article)
}

// GetArticleByID mengambil artikel berdasarkan ID
func (s *articleService) GetArticleByID(id int) (*model.Article, error) {
	return s.repo.GetArticleByID(id) // Memanggil repositori untuk mendapatkan artikel berdasarkan ID
}

// DeleteArticle menghapus artikel berdasarkan ID
func (s *articleService) DeleteArticle(id int) error {
	return s.repo.DeleteArticle(id) // Memanggil repositori untuk menghapus artikel berdasarkan ID
}

// GetAllArticles mengambil semua artikel
func (s *articleService) GetAllArticles() ([]model.Article, error) {
	return s.repo.GetAllArticles() // Memanggil repositori untuk mendapatkan semua artikel
}
