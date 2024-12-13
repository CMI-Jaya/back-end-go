package service

import (
	"database/sql"
	"errors"
	"go-project/internal/admin/model"
	"go-project/internal/admin/repository"
	"go-project/pkg/utils"
)

// Fungsi untuk autentikasi admin
func AuthenticateAdmin(db *sql.DB, email, password string) (model.User, string, error) {
	// Ambil data admin dari database
	admin, err := repository.GetAdminByEmail(db, email)
	if err != nil {
		return admin, "", err
	}

	// Cek password yang dimasukkan dengan password yang ada di database
	if !utils.CheckPasswordHash(password, admin.Password) {
		return admin, "", errors.New("invalid credentials")
	}

	// Generate token jika password valid
	token, err := utils.GenerateJWT(admin)
	if err != nil {
		return admin, "", err
	}

	return admin, token, nil
}
