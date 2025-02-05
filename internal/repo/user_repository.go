package repo

import (
	"errors"
	"mine/internal/model"
)

type UserRepository struct {
	users map[string]model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]model.User)}
}

func (r *UserRepository) Save(user model.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return errors.New("missing required fields")
	}
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(id string) (model.User, error) {
	user, ok := r.users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) Update(user model.User) error {
	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}
