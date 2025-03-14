package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"onboarding-go/internal/model"
	"onboarding-go/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("saving user")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.SaveUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *UserHandler) FindUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	slog.Info(fmt.Sprintf("finding user with id: %s", id))
	user, err := h.service.FindUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
