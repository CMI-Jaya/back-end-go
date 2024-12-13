package model

import (
	"time"
)

// Article represents the structure of an article.
type Article struct {
	ID              int       `json:"id"`
	CategoryID      int       `json:"category_id"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	Tags            []string  `json:"tags"`
	Content         string    `json:"content"`
	Message         string    `json:"message"`
	Thumbnail       string    `json:"thumbnail"`
	AltThumbnail    string    `json:"alt_thumbnail"`
	Banner          string    `json:"banner"`
	AltBanner       string    `json:"alt_banner"`
	Poster          string    `json:"poster"`
	AltPoster       string    `json:"alt_poster"`
	LinkVideo       string    `json:"link_video"`
	Status          string    `json:"status"`
	MetaTitle       string    `json:"meta_title"`
	MetaDescription string    `json:"meta_description"`
	AuthorID        int       `json:"author_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
