package repo

import (
	"context"
	"errors"
	"github.com/camilasimoess/onboarding-go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(user *model.User) error
	FindByID(id string) (*model.User, error)
	FindByNameAndLastName(firstName, lastName string) (*model.User, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database, collectionName string) *MongoUserRepository {
	collection := database.Collection(collectionName)
	return &MongoUserRepository{collection: collection}
}

func (r *MongoUserRepository) Save(user *model.User) error {
	user.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *MongoUserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindByNameAndLastName(firstName, lastName string) (*model.User, error) {
	var user model.User
	filter := bson.M{"firstname": firstName, "lastname": lastName}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
