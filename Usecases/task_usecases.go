package usecases

import (
	"fmt"

	domain "github.com/zaahidali/task_manager_api/Domain"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Implements the use cases related to tasks, such as creating, updating, retrieving, and deleting tasks.

// Get all tasks

type TaskUseCase struct {
	TaskRepository repositories.TaskRepositoryInterface
	UserRepository repositories.UserRepositoryInterface
}

func (taskuseCase *TaskUseCase) GetAlltasks() (datas []domain.Task, err error) {
	result, err := taskuseCase.TaskRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Get specific task

func (taskuseCase *TaskUseCase) GetSpecificTask(id primitive.ObjectID) (datas domain.Task, err error) {
	result, err := taskuseCase.TaskRepository.GetSpecificTask(id)
	if err != nil {
		return domain.Task{}, err
	}
	return result, nil
}

// Create task
func (taskuseCase *TaskUseCase) CreateTask(task domain.Task) (interface{}, error) {
	result, err := taskuseCase.TaskRepository.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// Update task
func (taskuseCase *TaskUseCase) UpdateTask(id primitive.ObjectID, task domain.Task) (interface{}, error) {
	result, err := taskuseCase.TaskRepository.UpdateTask(id, task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete task
func (taskuseCase *TaskUseCase) DeleteTask(id primitive.ObjectID) (interface{}, error) {
	result, err := taskuseCase.TaskRepository.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("task not found")
	}
	return result.DeletedCount, nil
}
