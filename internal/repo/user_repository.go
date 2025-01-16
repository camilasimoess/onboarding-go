package repo

import "mine/internal/model"

type UserRepository struct {
	users map[string]model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]model.User)}
}

func (r *UserRepository) Save(user model.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(id string) (model.User, error) {
	user, ok := r.users[id]
	if !ok {
		return model.User{}, nil
	}
	return user, nil
}
