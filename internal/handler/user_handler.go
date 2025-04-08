package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/camilasimoess/onboarding-go/internal/model"
	"github.com/camilasimoess/onboarding-go/internal/service"
	"log/slog"
	"net/http"
	"regexp"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	slog.Info("saving user")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	emailR := regexp.MustCompile(emailRegex)
	if !emailR.MatchString(user.Email) {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}

	err = h.service.CreateUser(ctx, &user)
	if err != nil {
		var validationError *service.ValidationError
		if errors.As(err, &validationError) {
			http.Error(w, validationError.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	slog.Info(fmt.Sprintf("finding user with id: %s", id))
	user, err := h.service.GetUser(ctx, id)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
