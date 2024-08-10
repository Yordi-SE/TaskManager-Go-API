package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// create user
func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := models.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// find user
func FindUser(user_id primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var user models.User
	err := models.UserCollection.FindOne(ctx, bson.M{"_id": user_id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func Promote(user_id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	r, err := models.UserCollection.UpdateByID(ctx, user_id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "role", Value: "admin"},
			},
		},
	})
	if r.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return err
	}
	return nil

}

// find user by user_name
func FindUserByName(user_name string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var user models.User
	err := models.UserCollection.FindOne(ctx, bson.M{"username": user_name}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
