package handler

import (
	"context"
	"encoding/json"
	"errors"
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

func validateRequest(user model.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return errors.New("missing required fields")
	}

	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	emailR := regexp.MustCompile(emailRegex)
	if !emailR.MatchString(user.Email) {
		return errors.New("invalid email")
	}
	return nil
}

func (h *UserHandler) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	slog.Info("request to create user received")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		slog.Error("failed to decode request body", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateRequest(user)
	if err != nil {
		slog.Error("validation failed", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateUser(ctx, &user)
	if err != nil {
		var validationError *service.ValidationError
		if errors.As(err, &validationError) {
			slog.Error("service validation failed", slog.String("error", validationError.Error()))
			http.Error(w, validationError.Error(), http.StatusBadRequest)
			return
		}
		slog.Error("failed to create user", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("user created successfully", slog.Any("user", user))
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	slog.Info("received request to get user", slog.String("id", id))
	user, err := h.service.GetUser(ctx, id)

	if err != nil {
		slog.Error("failed to get user", slog.String("error", err.Error()))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if user == nil {
		slog.Error("user not found", slog.String("id", id))
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	slog.Info("user found", slog.Any("user", user))
	json.NewEncoder(w).Encode(user)
}
