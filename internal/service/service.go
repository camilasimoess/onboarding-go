package service

import (
	"github.com/camilasimoess/onboarding-go/internal/model"
	"github.com/camilasimoess/onboarding-go/internal/repo"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	if err := s.validateUser(*user); err != nil {
		return err
	}
	return s.repo.Save(user)
}

func (s *UserService) GetUser(id string) (*model.User, error) {
	return s.repo.FindByID(id)
}
