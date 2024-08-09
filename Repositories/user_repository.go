package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// create user
func CreateUser(ctx *gin.Context, user models.User) (*mongo.InsertOneResult, error) {
	result, err := models.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// find user
func FindUser(ctx *gin.Context, user_id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := models.UserCollection.FindOne(ctx, bson.M{"_id": user_id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func Promote(ctx *gin.Context, user_id primitive.ObjectID) error {
	_, err := models.UserCollection.UpdateByID(ctx, user_id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "role", Value: "admin"},
			},
		},
	})
	if err != nil {
		return err
	}
	return nil

}
