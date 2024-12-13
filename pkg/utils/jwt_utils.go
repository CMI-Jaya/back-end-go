package utils

import (
	"go-project/internal/admin/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key untuk signing JWT
var jwtKey = []byte("your_secret_key")

// Membuat token untuk admin
func GenerateJWT(admin model.User) (string, error) {
	// Membuat klaim (claims) untuk JWT
	claims := &jwt.StandardClaims{
		Subject:   admin.Email,
		Issuer:    "admin_app",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	// Membuat token dengan signing method HMAC dan klaim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Validasi JWT dan mendapatkan klaim
func ValidateJWT(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	// Mendapatkan klaim
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		return claims, nil
	}
	return nil, err
}
