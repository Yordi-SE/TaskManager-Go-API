package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/Infrastructure"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user usecase

func Register(ctx *gin.Context, user models.User) (interface{}, error) {

	hashedPassword, err := Infrastructure.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	userCount, err := repositories.Count(ctx, models.UserCollection)
	if err != nil {
		return nil, err
	}
	if userCount == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	result, errs := repositories.CreateUser(ctx, user)
	if errs != nil {
		return nil, errs
	}
	return result.InsertedID, nil
}

// promote user
func Promote(ctx *gin.Context, user_id primitive.ObjectID) error {
	err := repositories.Promote(ctx, user_id)
	if err != nil {
		return err
	}
	return nil
}
