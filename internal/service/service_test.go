package service

import (
	"errors"
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

func Test_FindByID(t *testing.T) {
	mockRepo := repo.NewMockery_UserRepository(t)

	t.Run("Positive - User Found", func(t *testing.T) {
		expectedUser := &model.User{
			ID:        "67dd9bdf4a516d5f0bd9f8a4",
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}

		mockRepo.EXPECT().FindByID(expectedUser.ID).Return(expectedUser, nil)

		user, err := mockRepo.FindByID(expectedUser.ID)
		assert.Nil(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Negative - User Not Found", func(t *testing.T) {
		nonExistentID := "non existent id"

		mockRepo.EXPECT().FindByID(nonExistentID).Return(nil, nil)

		user, err := mockRepo.FindByID(nonExistentID)
		assert.Nil(t, user)
		assert.Nil(t, err)
	})

	t.Run("Negative - Error Occurred", func(t *testing.T) {
		errorID := "error id"
		expectedError := errors.New("any some error")

		mockRepo.EXPECT().FindByID(errorID).Return(nil, expectedError)

		user, err := mockRepo.FindByID(errorID)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}
