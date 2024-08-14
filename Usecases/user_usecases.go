package usecases

import (
	"fmt"

	domain "github.com/zaahidali/task_manager_api/Domain"
	"github.com/zaahidali/task_manager_api/Infrastructure"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUseCase struct {
	TaskRepository repositories.TaskRepositoryInterface
	UserRepository repositories.UserRepositoryInterface
	Infrastracture Infrastructure.InfrastructureInterface
}

// user usecase

func (userusecases *UserUseCase) Register(user domain.User) (interface{}, error) {

	hashedPassword, err := userusecases.Infrastracture.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	userCount, err := userusecases.TaskRepository.Count(repositories.UserCollection)
	if err != nil {
		return nil, err
	}
	if userCount == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	var existing domain.User
	existing, errs := userusecases.UserRepository.FindUserByName(user.UserName)
	if errs != nil {
		if errs != mongo.ErrNoDocuments {
			return nil, errs
		}
	}
	if existing != (domain.User{}) && existing.UserName == user.UserName {
		return nil, fmt.Errorf("user already exists")
	}
	if existing.UserName == user.UserName {
		return nil, fmt.Errorf("user already exists")
	}
	result, errs := userusecases.UserRepository.CreateUser(user)
	if errs != nil {
		return nil, errs
	}
	return result.InsertedID, nil
}

// promote user
func (userusecases *UserUseCase) Promote(user_id primitive.ObjectID) error {
	result, err := userusecases.UserRepository.Promote(user_id)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("user is admin")
	}
	return nil
}

func (userusecases *UserUseCase) Login(data domain.User) (string, error) {
	var result domain.User
	result, errs := userusecases.UserRepository.FindUserByName(data.UserName)
	if errs != nil {
		return "", errs
	}
	err := userusecases.Infrastracture.ComparePasswords(result.Password, data.Password)
	if err != nil {
		return "", err
	}
	jwtToken, err := userusecases.Infrastracture.GenerateToken(result.UserName, result.ID, result.Role)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
