package repositories

import (
	"context"
	"time"

	domain "github.com/zaahidali/task_manager_api/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
}

// Abstracts the data access logic.
// The repository pattern is a design pattern that isolates the data access logic from the rest of the application.
type UserRepositoryInterface interface {
	CreateUser(user domain.User) (*mongo.InsertOneResult, error)
	FindUser(user_id primitive.ObjectID) (domain.User, error)
	FindUserByName(user_name string) (domain.User, error)
	Promote(user_id primitive.ObjectID) (*mongo.UpdateResult, error)
}

// create user
func (*UserRepository) CreateUser(user domain.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// find user
func (*UserRepository) FindUser(user_id primitive.ObjectID) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var user domain.User
	err := UserCollection.FindOne(ctx, bson.M{"_id": user_id}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (*UserRepository) Promote(user_id primitive.ObjectID) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	r, err := UserCollection.UpdateByID(ctx, user_id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "role", Value: "admin"},
			},
		},
	})

	if err != nil {
		return nil, err
	}
	return r, nil

}

// find user by user_name
func (*UserRepository) FindUserByName(user_name string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var user domain.User
	err := UserCollection.FindOne(ctx, bson.M{"username": user_name}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
