package service

import (
	"mine/internal/model"
	"mine/internal/repo"
)

type UserService struct {
	repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) FindUserByID(id string) (model.User, error) {
	return s.UserRepository.FindByID(id)
}

func (s *UserService) SaveUser(user model.User) error {
	return s.UserRepository.Save(user)
}
