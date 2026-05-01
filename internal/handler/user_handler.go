package handler

import (
	"encoding/json"
	"go-ticket/internal/config"
	"go-ticket/internal/domain"
	"go-ticket/internal/dto"
	"mime/multipart"
	"net/http"
	"strings"
)

type UserHandler struct {
	userService domain.UserService
	cfg         *config.Config
}

func NewUserHandler(userService domain.UserService, cfg *config.Config) *UserHandler {
	return &UserHandler{
		userService: userService,
		cfg:         cfg,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.userService.GetProfileByID(userID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	if user == nil {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	writeJSON(w, http.StatusOK, h.toUserResponse(user))
}

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var input domain.UpdateProfileInput
	var fileHeader *multipart.FileHeader

	contentType := r.Header.Get("Content-Type")

	if strings.HasPrefix(contentType, "multipart/form-data") {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			writeError(w, http.StatusBadRequest, "invalid form data")
			return
		}

		if name := r.FormValue("name"); name != "" {
			input.Name = &name
		}

		if phone := r.FormValue("phone"); phone != "" {
			input.Phone = &phone
		}

		_, header, err := r.FormFile("profile")
		if err == nil {
			fileHeader = header
		}

	} else {
		var req dto.UpdateProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid request")
			return
		}

		if req.Name != "" {
			input.Name = &req.Name
		}
		if req.Phone != "" {
			input.Phone = &req.Phone
		}
	}

	user, err := h.userService.UpdateProfile(userID, input, fileHeader)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, h.toUserResponse(user))
}

func (h *UserHandler) buildProfileURL(profile string) string {
	if profile == "" {
		return ""
	}
	if strings.HasPrefix(profile, "http") {
		return profile
	}
	if strings.HasPrefix(profile, "uploads/") {
		return h.cfg.BaseURL + "/" + profile
	}
	return h.cfg.BaseURL + "/uploads/" + profile
}

func (h *UserHandler) toUserResponse(user *domain.User) dto.UserResponse {
	if user == nil {
		return dto.UserResponse{}
	}

	return dto.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Role:    user.Role,
		Profile: h.buildProfileURL(user.Profile),
	}
}
