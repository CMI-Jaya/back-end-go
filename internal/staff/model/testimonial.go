package model

import "time"

type Testimonial struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Comment      string    `json:"comment"`
	PhotoProfile string    `json:"photo_profile"`
	CategoryID   int       `json:"category_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
