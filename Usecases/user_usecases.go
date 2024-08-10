package usecases

import (
	"fmt"

	"github.com/zaahidali/task_manager_api/Infrastructure"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// user usecase

func Register(user models.User) (interface{}, error) {

	hashedPassword, err := Infrastructure.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	userCount, err := repositories.Count(models.UserCollection)
	if err != nil {
		return nil, err
	}
	if userCount == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	var existing models.User
	existing, errs := repositories.FindUserByName(user.UserName)
	if errs != nil {
		if errs != mongo.ErrNoDocuments {
			return nil, errs
		}
	}
	if existing != (models.User{}) && existing.UserName == user.UserName {
		return nil, fmt.Errorf("user already exists")
	}
	if existing.UserName == user.UserName {
		return nil, fmt.Errorf("user already exists")
	}
	result, errs := repositories.CreateUser(user)
	if errs != nil {
		return nil, errs
	}
	return result.InsertedID, nil
}

// promote user
func Promote(user_id primitive.ObjectID) error {
	err := repositories.Promote(user_id)
	if err != nil {
		return err
	}
	return nil
}

func Login(data models.User) (string, error) {
	var result models.User
	result, errs := repositories.FindUserByName(data.UserName)
	if errs != nil {
		return "", errs
	}
	err := Infrastructure.ComparePasswords(result.Password, data.Password)
	fmt.Println(result.Password, data.Password, "<=password")
	if err != nil {
		return "", err
	}
	jwtToken, err := Infrastructure.GenerateToken(result.UserName, result.ID, result.Role)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
