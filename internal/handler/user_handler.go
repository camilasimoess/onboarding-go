package handler

import (
	"encoding/json"
	"net/http"
	"onboarding-go/internal/model"
	"onboarding-go/internal/service"
	"strings"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func (h *UserHandler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/find/")
	user, err := h.service.FindUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
