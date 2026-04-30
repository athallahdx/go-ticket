package handler

import (
	"go-ticket/internal/config"
	"go-ticket/internal/domain"
)

type AdminUserHandler struct {
	userService domain.UserService
	cfg         *config.Config
}

func NewAdminUserHandler(userService domain.UserService, cfg *config.Config) *AdminUserHandler {
	return &AdminUserHandler{
		userService: userService,
		cfg:         cfg,
	}
}
