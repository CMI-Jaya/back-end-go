package model

// Struktur untuk model Notification
type Notification struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Status    string `json:"status"` // unread, read
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
