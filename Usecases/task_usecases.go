package usecases

import (
	domain "github.com/zaahidali/task_manager_api/Domain"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Implements the use cases related to tasks, such as creating, updating, retrieving, and deleting tasks.

// Get all tasks
func GetAlltasks() (datas []domain.Task, err error) {
	result, err := repositories.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Get specific task

func GetSpecificTask(id primitive.ObjectID) (datas domain.Task, err error) {
	result, err := repositories.GetSpecificTask(id)
	if err != nil {
		return domain.Task{}, err
	}
	return result, nil
}

// Create task
func CreateTask(task domain.Task) (interface{}, error) {
	result, err := repositories.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// Update task
func UpdateTask(id primitive.ObjectID, task domain.Task) (interface{}, error) {
	result, err := repositories.UpdateTask(id, task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete task
func DeleteTask(id primitive.ObjectID) (interface{}, error) {
	result, err := repositories.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	return result.DeletedCount, nil
}
