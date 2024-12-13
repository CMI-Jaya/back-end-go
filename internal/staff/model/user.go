package model

import "time"

// User represents a user in the system, such as staff or admin.
type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`   // Role can be "admin", "staff", or "user"
	Status      string    `json:"status"` // "active", "inactive"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
