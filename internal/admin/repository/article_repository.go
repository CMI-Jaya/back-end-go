package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"go-project/internal/admin/model"
	"log"
)

// ArticleRepository adalah interface yang mendefinisikan metode-metode untuk berinteraksi dengan data artikel di database.
type ArticleRepository interface {
	CreateArticle(article *model.Article) error    // Menyimpan artikel baru ke database
	GetArticleByID(id int) (*model.Article, error) // Mengambil artikel berdasarkan ID
	UpdateArticle(article *model.Article) error    // Memperbarui data artikel yang sudah ada
	DeleteArticle(id int) error                    // Menghapus artikel berdasarkan ID
	GetAllArticles() ([]model.Article, error)      // Mengambil semua artikel
}

// articleRepository adalah implementasi dari ArticleRepository, menyimpan koneksi ke database.
type articleRepository struct {
	db *sql.DB
}

// NewArticleRepository adalah konstruktor yang mengembalikan instance baru dari articleRepository dengan koneksi database yang diberikan.
func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{db: db}
}

// validStatus memeriksa apakah status artikel yang diberikan valid (di antara status yang sudah ditentukan).
func validStatus(status string) bool {
	validStatuses := []string{"approval", "pending approval", "rejected"}
	for _, s := range validStatuses {
		if status == s {
			return true
		}
	}
	return false
}

// UpdateArticleStatus memperbarui status artikel berdasarkan ID artikel dan status yang diberikan.
// Jika status tidak valid, maka akan mengembalikan error.
func (r *articleRepository) UpdateArticleStatus(id int, status string) error {
	if !validStatus(status) {
		return errors.New("invalid status")
	}

	query := `UPDATE articles SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}

// CreateArticle menyimpan artikel baru ke dalam database. Jika kategori yang diberikan tidak ada, akan mengembalikan error.
func (r *articleRepository) CreateArticle(article *model.Article) error {
	var categoryExists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1)", article.CategoryID).Scan(&categoryExists)
	if err != nil {
		log.Printf("Error checking category existence: %v", err)
		return err
	}

	if !categoryExists {
		return errors.New("invalid category_id: category does not exist")
	}

	query := `
        INSERT INTO articles (
            category_id, title, slug, tags, content, message, thumbnail, alt_thumbnail, banner, 
            alt_banner, poster, alt_poster, link_video, status, meta_title, meta_description, 
            author_id, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
        )
    `
	_, err = r.db.Exec(query, article.CategoryID, article.Title, article.Slug, article.Tags, article.Content,
		article.Message, article.Thumbnail, article.AltThumbnail, article.Banner, article.AltBanner, article.Poster,
		article.AltPoster, article.LinkVideo, article.Status, article.MetaTitle, article.MetaDescription, article.AuthorID,
		article.CreatedAt, article.UpdatedAt)

	if err != nil {
		log.Printf("Error creating article: %v", err)
		return err
	}

	return nil
}

// GetArticleByID mengambil artikel berdasarkan ID dari database. Jika artikel tidak ditemukan, mengembalikan error.
func (r *articleRepository) GetArticleByID(id int) (*model.Article, error) {
	query := `SELECT * FROM articles WHERE id = $1`
	row := r.db.QueryRow(query, id)
	var article model.Article

	// Menangani tags sebagai JSON
	var tags json.RawMessage
	err := row.Scan(
		&article.ID, &article.CategoryID, &article.Title, &article.Slug,
		&tags, // Menyimpan dalam json.RawMessage
		&article.Content, &article.Message, &article.Thumbnail,
		&article.AltThumbnail, &article.Banner, &article.AltBanner,
		&article.Poster, &article.AltPoster, &article.LinkVideo,
		&article.Status, &article.MetaTitle, &article.MetaDescription,
		&article.AuthorID, &article.CreatedAt, &article.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("article with id %d not found", id)
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	// Jika tags valid dalam format JSON, lakukan unmarshal ke dalam slice string
	if len(tags) > 0 {
		var tagList []string
		if err := json.Unmarshal(tags, &tagList); err != nil {
			return nil, fmt.Errorf("error unmarshaling tags JSON: %v", err)
		}
		article.Tags = tagList
	} else {
		article.Tags = []string{}
	}

	return &article, nil
}

// UpdateArticle memperbarui data artikel yang sudah ada berdasarkan artikel yang diberikan.
func (r *articleRepository) UpdateArticle(article *model.Article) error {
	query := `
		UPDATE articles SET
			category_id = $1, title = $2, slug = $3, tags = $4, content = $5, message = $6, thumbnail = $7,
			alt_thumbnail = $8, banner = $9, alt_banner = $10, poster = $11, alt_poster = $12, link_video = $13,
			status = $14, meta_title = $15, meta_description = $16, author_id = $17, updated_at = $18
		WHERE id = $19
	`
	_, err := r.db.Exec(query, article.CategoryID, article.Title, article.Slug, article.Tags, article.Content,
		article.Message, article.Thumbnail, article.AltThumbnail, article.Banner, article.AltBanner, article.Poster,
		article.AltPoster, article.LinkVideo, article.Status, article.MetaTitle, article.MetaDescription,
		article.AuthorID, article.UpdatedAt, article.ID)

	if err != nil {
		log.Printf("Error updating article: %v", err)
		return err
	}

	return nil
}

// DeleteArticle menghapus artikel berdasarkan ID dari database.
func (r *articleRepository) DeleteArticle(id int) error {
	query := `DELETE FROM articles WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting article: %v", err)
		return err
	}
	return nil
}

// GetAllArticles mengambil semua artikel dari database dan mengembalikannya dalam bentuk slice dari model.Article.
func (r *articleRepository) GetAllArticles() ([]model.Article, error) {
	query := `SELECT * FROM articles`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error retrieving articles: %v", err)
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		var tagsData []byte // Scan data byte mentah terlebih dahulu
		if err := rows.Scan(
			&article.ID, &article.CategoryID, &article.Title, &article.Slug, &tagsData, &article.Content,
			&article.Message, &article.Thumbnail, &article.AltThumbnail, &article.Banner, &article.AltBanner,
			&article.Poster, &article.AltPoster, &article.LinkVideo, &article.Status, &article.MetaTitle,
			&article.MetaDescription, &article.AuthorID, &article.CreatedAt, &article.UpdatedAt,
		); err != nil {
			log.Printf("Error scanning article: %v", err)
			continue
		}

		// Jika tagsData berisi array yang diserialisasi, kita bisa melakukan unmarshal ke dalam []string
		if err := json.Unmarshal(tagsData, &article.Tags); err != nil {
			log.Printf("Error unmarshaling tags data: %v", err)
		}

		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return articles, nil
}
