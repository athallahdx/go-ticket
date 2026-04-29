// domain/auth.go  ← add this new file
package domain

import "go-ticket/internal/pkg/jwt"

type AuthService interface {
	Register(name, email, password string) (*User, error)
	Login(email, password string) (string, error)
	GetProfile(userID int64) (*User, error)
	ChangePassword(userID int64, oldPassword, newPassword string) error
	ValidateToken(token string) (*jwt.Claims, error)
}
