package handler

import (
	"github.com/camilasimoess/onboarding-go/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateRequest(t *testing.T) {
	t.Run("Positive - Valid User", func(t *testing.T) {
		user := model.User{
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}

		err := validateRequest(user)
		assert.Nil(t, err)
	})

	t.Run("Negative - Invalid Email", func(t *testing.T) {
		user := model.User{
			FirstName: "Camila",
			LastName:  "Simoes",
			Email:     "camila.simoes",
			Age:       25,
		}

		err := validateRequest(user)
		assert.Equal(t, "invalid email", err.Error())
	})

	t.Run("Negative - Missing Required Fields", func(t *testing.T) {
		user := model.User{
			FirstName: "",
			LastName:  "Simoes",
			Email:     "camila.simoes@test.com",
			Age:       25,
		}

		err := validateRequest(user)
		assert.Equal(t, "missing required fields", err.Error())
	})
}
