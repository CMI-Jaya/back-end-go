package model

// User mewakili data pengguna admin di aplikasi.
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // Misalnya "admin" atau "superadmin"
}
