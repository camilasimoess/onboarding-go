package service

import (
	"context"
	"errors"
	"github.com/camilasimoess/onboarding-go/internal/model"
	"github.com/camilasimoess/onboarding-go/internal/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	ctx := context.Background()

	t.Run("Positive - Valid User", func(t *testing.T) {
		mockRepo := repo.NewMockery_UserRepository(t)
		userService := NewUserService(mockRepo)
		user := model.User{
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}
		mockRepo.EXPECT().FindByNameAndLastName(ctx, user.FirstName, user.LastName).Return(nil, nil)
		mockRepo.EXPECT().Save(ctx, &user).Return(nil)

		err := userService.CreateUser(ctx, &user)
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

		err := userService.CreateUser(ctx, &user)
		assert.Equal(t, ErrorInvalidAge, err)
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
		mockRepo.EXPECT().FindByNameAndLastName(ctx, user.FirstName, user.LastName).Return(&existingUser, nil)

		err := userService.CreateUser(ctx, &user)
		assert.Equal(t, ErrorUserAlreadyExists, err)
	})
}

func Test_FindByID(t *testing.T) {
	ctx := context.Background()
	mockRepo := repo.NewMockery_UserRepository(t)

	t.Run("Positive - User Found", func(t *testing.T) {
		expectedUser := &model.User{
			ID:        "67dd9bdf4a516d5f0bd9f8a4",
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}

		mockRepo.EXPECT().FindByID(ctx, expectedUser.ID).Return(expectedUser, nil)

		user, err := mockRepo.FindByID(ctx, expectedUser.ID)
		assert.Nil(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Negative - User Not Found", func(t *testing.T) {
		nonExistentID := "non existent id"

		mockRepo.EXPECT().FindByID(ctx, nonExistentID).Return(nil, nil)

		user, err := mockRepo.FindByID(ctx, nonExistentID)
		assert.Nil(t, user)
		assert.Nil(t, err)
	})

	t.Run("Negative - Error Occurred", func(t *testing.T) {
		errorID := "error id"
		expectedError := errors.New("any some error")

		mockRepo.EXPECT().FindByID(ctx, errorID).Return(nil, expectedError)

		user, err := mockRepo.FindByID(ctx, errorID)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}
