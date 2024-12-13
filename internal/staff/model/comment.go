package model

import "time"

type Comment struct {
	ID        int       `json:"id"`
	ArticleID int       `json:"article_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Comment   string    `json:"comment"`
	ParentID  *int      `json:"parent_id,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
