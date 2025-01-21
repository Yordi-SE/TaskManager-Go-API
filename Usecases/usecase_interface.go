package usecases

import (
	domain "github.com/zaahidali/task_manager_api/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCaseInterface interface {
	GetAlltasks() ([]domain.Task, error)
	GetSpecificTask(id primitive.ObjectID) (domain.Task, error)
	CreateTask(task domain.Task) (interface{}, error)
	UpdateTask(id primitive.ObjectID, task domain.Task) (interface{}, error)
	DeleteTask(id primitive.ObjectID) (interface{}, error)
}

type UserUseCaseInterface interface {
	Register(user domain.User) (interface{}, error)
	Login(user domain.User) (string, error)
	Promote(user_id primitive.ObjectID) error
}
