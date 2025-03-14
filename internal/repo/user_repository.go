package repo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"onboarding-go/internal/model"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database, dbName, collectionName string) *UserRepository {
	collection := database.Collection(collectionName)
	return &UserRepository{collection: collection}
}

func (r *UserRepository) Save(user *model.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return errors.New("missing required fields")
	}
	_, err := r.collection.InsertOne(nil, user)
	return err
}

func (r *UserRepository) FindByID(id string) (model.User, error) {
	var user model.User
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}
	return user, nil
}
