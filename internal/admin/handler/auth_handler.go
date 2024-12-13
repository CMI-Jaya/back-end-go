package handler

import (
	"encoding/json"
	"go-project/db"
	"go-project/internal/admin/model"
	"go-project/internal/admin/service"
	"go-project/pkg/utils"
	"log"
	"net/http"
)

// Fungsi untuk registrasi pengguna (dengan role dinamis)
func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi role
	validRoles := map[string]bool{"admin": true, "staff": true, "user": true}
	if !validRoles[user.Role] {
		http.Error(w, "Invalid role specified", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Simpan data pengguna ke database
	err = db.SaveUser(user.Email, hashedPassword, user.Role)
	if err != nil {
		http.Error(w, "Error saving user to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully")
}

// Fungsi untuk login pengguna
func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("Login endpoint hit")

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginData)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Println("Request body decoded:", loginData)

	// Check database connection
	dbConn := db.GetDB()
	if dbConn == nil {
		log.Println("Database connection error")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	log.Println("Database connection established")

	// Authenticate user
	user, token, err := service.AuthenticateAdmin(dbConn, loginData.Email, loginData.Password)
	if err != nil {
		log.Println("Authentication failed:", err)
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}
	log.Println("Authentication successful, user:", user)

	// Send response
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Fungsi untuk logout pengguna
func LogoutAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Del("Authorization")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User logged out successfully")
}
