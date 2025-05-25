package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for Swagger or health check
		if strings.HasPrefix(r.URL.Path, "/swagger") || r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
			return
		}

		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		secret := os.Getenv("JWT_SECRET")

		if tokenStr == "" || secret == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
