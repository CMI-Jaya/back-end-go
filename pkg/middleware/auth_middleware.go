package middleware

import (
	"context"
	"go-project/pkg/utils"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Validasi token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			log.Println("Invalid token:", err)
			return
		}

		// Menambahkan klaim ke context untuk digunakan di handler berikutnya
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
