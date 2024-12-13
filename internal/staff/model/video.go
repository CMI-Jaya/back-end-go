package model

import "time"

type Video struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	LinkVideo       string    `json:"link_video"`
	CategoryID      int       `json:"category_id"`
	Status          string    `json:"status"`    // Ini penting
	AuthorID        int       `json:"author_id"` // Ini penting
	MetaTitle       string    `json:"meta_title"`
	MetaDescription string    `json:"meta_description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
