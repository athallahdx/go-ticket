package middleware

import (
	"context"
	"net/http"
	"strings"

	"go-ticket/internal/config"
	myJwt "go-ticket/pkg/jwt"
	"go-ticket/pkg/response"
)

type contextKey string

const (
	ContextUserID contextKey = "user_id"
	ContextRole   contextKey = "role"
)

func AuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if !strings.HasPrefix(authHeader, "Bearer ") {
				response.Error(w, http.StatusUnauthorized, "missing or invalid token")
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := myJwt.ValidateToken(tokenString, cfg.JWTSecret)
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "invalid token")
				return
			}

			if myJwt.IsExpired(claims) {
				response.Error(w, http.StatusUnauthorized, "token expired")
				return
			}

			ctx := context.WithValue(r.Context(), ContextUserID, claims.UserID)
			ctx = context.WithValue(ctx, ContextRole, claims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
