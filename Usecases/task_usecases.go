package usecases

import (
	"github.com/gin-gonic/gin"
	domain "github.com/zaahidali/task_manager_api/Domain"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Implements the use cases related to tasks, such as creating, updating, retrieving, and deleting tasks.

// Get all tasks
func GetAlltasks(ctx *gin.Context) (datas []domain.Task, err error) {
	result, err := repositories.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Get specific task

func GetSpecificTask(ctx *gin.Context, id primitive.ObjectID) (datas domain.Task, err error) {
	result, err := repositories.GetSpecificTask(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}
	return result, nil
}

// Create task
func CreateTask(ctx *gin.Context, task domain.Task) (interface{}, error) {
	result, err := repositories.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// Update task
func UpdateTask(ctx *gin.Context, id primitive.ObjectID, task domain.Task) (interface{}, error) {
	result, err := repositories.UpdateTask(ctx, id, task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete task
func DeleteTask(ctx *gin.Context, id primitive.ObjectID) (interface{}, error) {
	result, err := repositories.DeleteTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return result.DeletedCount, nil
}
