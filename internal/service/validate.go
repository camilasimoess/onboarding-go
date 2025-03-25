package service

import (
	"github.com/camilasimoess/onboarding-go/internal/model"
	"regexp"
)

var (
	ErrorInvalidAge            = ValidationError{Message: "user must be at least 18 years old"}
	ErrorInvalidEmail          = ValidationError{Message: "invalid email"}
	ErrorMissingRequiredFields = ValidationError{Message: "missing required fields"}
	ErrorUserAlreadyExists     = ValidationError{Message: "user already exists"}
)

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func (s *UserService) validateUser(user model.User) error {
	existingUser, err := s.repo.FindByNameAndLastName(user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return ErrorUserAlreadyExists
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return ErrorMissingRequiredFields
	}

	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	emailR := regexp.MustCompile(emailRegex)
	if !emailR.MatchString(user.Email) {
		return ErrorInvalidEmail
	}

	if user.Age < 18 {
		return ErrorInvalidAge
	}
	return nil
}
