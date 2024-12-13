package repository

import (
	"database/sql"
	"go-project/internal/admin/model"
	"log"
)

// Fungsi untuk mengambil data admin berdasarkan email
func GetAdminByEmail(db *sql.DB, email string) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, email, password, role FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil // Tidak ada user ditemukan
		}
		log.Fatal(err)
	}
	return user, err
}
