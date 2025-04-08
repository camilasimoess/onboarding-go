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
	Save(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByNameAndLastName(ctx context.Context, firstName, lastName string) (*model.User, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{collection: database.Collection("users")}
}

func (r *MongoUserRepository) Save(ctx context.Context, user *model.User) error {
	user.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindByNameAndLastName(ctx context.Context, firstName, lastName string) (*model.User, error) {
	var user model.User
	filter := bson.M{"firstname": firstName, "lastname": lastName}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
