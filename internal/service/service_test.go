package service

import (
	"github.com/camilasimoess/onboarding-go/internal/model"
	"github.com/camilasimoess/onboarding-go/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateUser(t *testing.T) {

	t.Run("Positive - Valid User", func(t *testing.T) {
		mockRepo := repo.NewMockery_UserRepository(t)
		userService := NewUserService(mockRepo)
		user := model.User{
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}
		mockRepo.EXPECT().FindByNameAndLastName(user.FirstName, user.LastName).Return(nil, nil)
		mockRepo.EXPECT().Save(&user).Return(nil)

		err := userService.CreateUser(&user)
		assert.Nil(t, err)
	})

	t.Run("Negative - Invalid Age", func(t *testing.T) {
		mockRepo := repo.NewMockery_UserRepository(t)
		userService := NewUserService(mockRepo)
		user := model.User{
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       17,
		}

		mockRepo.EXPECT().FindByNameAndLastName(user.FirstName, user.LastName).Return(nil, nil)

		err := userService.CreateUser(&user)
		assert.Equal(t, ErrorInvalidAge, err)
	})

	t.Run("Negative - Invalid Email", func(t *testing.T) {
		mockRepo := repo.NewMockery_UserRepository(t)
		userService := NewUserService(mockRepo)
		user := model.User{
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes",
			Age:       25,
		}

		mockRepo.EXPECT().FindByNameAndLastName(user.FirstName, user.LastName).Return(nil, nil)

		err := userService.CreateUser(&user)
		assert.Equal(t, ErrorInvalidEmail, err)
	})

	t.Run("Negative - Missing Required Fields", func(t *testing.T) {
		mockRepo := repo.NewMockery_UserRepository(t)
		userService := NewUserService(mockRepo)
		user := model.User{
			FirstName: "",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}

		mockRepo.EXPECT().FindByNameAndLastName(user.FirstName, user.LastName).Return(nil, nil)

		err := userService.CreateUser(&user)
		assert.Equal(t, ErrorMissingRequiredFields, err)
	})

	t.Run("Negative - User Already Exists", func(t *testing.T) {
		mockRepo := repo.NewMockery_UserRepository(t)
		userService := NewUserService(mockRepo)
		existingUser := model.User{
			ID:        "67dd9bdf4a516d5f0bd9f8a4",
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}

		user := model.User{
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}
		mockRepo.EXPECT().FindByNameAndLastName(user.FirstName, user.LastName).Return(&existingUser, nil)

		err := userService.CreateUser(&user)
		assert.Equal(t, ErrorUserAlreadyExists, err)
	})
}
