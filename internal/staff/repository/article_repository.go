package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-project/internal/staff/model"
)

type ArticleRepository struct {
	DB *sql.DB
}

func (r *ArticleRepository) SaveArticle(article model.Article) error {
	query := `
        INSERT INTO articles (
            category_id, title, slug, tags, content, message, thumbnail, alt_thumbnail, banner, 
            alt_banner, poster, alt_poster, link_video, status, meta_title, meta_description, 
            author_id, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
        )
    `
	_, err := r.DB.Exec(query, article.CategoryID, article.Title, article.Slug, article.Tags, article.Content,
		article.Message, article.Thumbnail, article.AltThumbnail, article.Banner, article.AltBanner, article.Poster,
		article.AltPoster, article.LinkVideo, article.Status, article.MetaTitle, article.MetaDescription, article.AuthorID,
		article.CreatedAt, article.UpdatedAt)
	return err
}

func (r *ArticleRepository) GetArticleByID(id int) (*model.Article, error) {
	query := `SELECT * FROM articles WHERE id = $1`
	row := r.DB.QueryRow(query, id)
	var article model.Article

	// Handle tags as JSON
	var tags json.RawMessage
	err := row.Scan(
		&article.ID, &article.CategoryID, &article.Title, &article.Slug,
		&tags, // Scanning into a json.RawMessage
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

	// If tags are valid JSON, unmarshal it into a string slice
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

func (r *ArticleRepository) GetAllArticles() ([]model.Article, error) {
	query := `SELECT id, title, content, category_id, status, author_id, meta_title, meta_description, created_at, updated_at FROM articles`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		if err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.CategoryID,
			&article.Status,
			&article.AuthorID,
			&article.MetaTitle,
			&article.MetaDescription,
			&article.CreatedAt,
			&article.UpdatedAt,
		); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
