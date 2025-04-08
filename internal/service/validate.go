package service

import (
	"context"
	"github.com/camilasimoess/onboarding-go/internal/model"
)

var (
	ErrorInvalidAge        = &ValidationError{Message: "user must be at least 18 years old"}
	ErrorUserAlreadyExists = &ValidationError{Message: "user already exists"}
)

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func (s *UserService) validateUser(ctx context.Context, user model.User) error {
	if user.Age < 18 {
		return ErrorInvalidAge
	}

	existingUser, err := s.repo.FindByNameAndLastName(ctx, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return ErrorUserAlreadyExists
	}
	return nil
}
