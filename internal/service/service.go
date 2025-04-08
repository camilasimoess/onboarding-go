package service

import (
	"context"
	"github.com/camilasimoess/onboarding-go/internal/model"
	"github.com/camilasimoess/onboarding-go/internal/repo"
	"log/slog"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	if err := s.validateUser(ctx, *user); err != nil {
		slog.Error("validation error", "error", err)
		return err
	}
	return s.repo.Save(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}
