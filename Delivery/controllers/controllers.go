package controllers

import (
	"fmt"

	usecases "github.com/zaahidali/task_manager_api/Usecases"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	domain "github.com/zaahidali/task_manager_api/Domain"
)

type TaskHandler struct {
	UserUsecase usecases.UserUseCaseInterface
	TaskUsecase usecases.TaskUseCaseInterface
}

func (taskcontroller *TaskHandler) GetTasks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := taskcontroller.TaskUsecase.GetAlltasks()
		if err != nil {
			ctx.IndentedJSON(500, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(200, result)
	}
}

func (taskcontroller *TaskHandler) GetTasksId() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		Id, errs := primitive.ObjectIDFromHex(id)
		if errs != nil {
			ctx.IndentedJSON(404, gin.H{"message": errs.Error()})
			return
		}
		tasks, err := taskcontroller.TaskUsecase.GetSpecificTask(Id)

		if err != nil {
			ctx.IndentedJSON(404, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(200, tasks)

	}
}

func (taskcontroller *TaskHandler) CreateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task domain.Task
		err := ctx.BindJSON(&task)
		if err != nil {
			ctx.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		}
		result, err := taskcontroller.TaskUsecase.CreateTask(task)
		if err != nil {
			fmt.Println(err)
			ctx.IndentedJSON(500, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(201, result)
	}
}

func (taskcontroller *TaskHandler) UpdateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		Id, errss := primitive.ObjectIDFromHex(id)
		if errss != nil {
			ctx.IndentedJSON(404, gin.H{"message": errss.Error()})
			return
		}
		var task domain.Task
		errs := ctx.BindJSON(&task)
		if errs != nil {
			ctx.IndentedJSON(400, gin.H{"message": errs.Error()})
			return
		}
		result, err := taskcontroller.TaskUsecase.UpdateTask(Id, task)

		if err != nil {
			ctx.IndentedJSON(500, gin.H{"message": err.Error()})
			return
		}

		ctx.IndentedJSON(200, result)
	}
}

func (taskcontroller *TaskHandler) DeleteTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		Id, errss := primitive.ObjectIDFromHex(id)
		if errss != nil {
			ctx.IndentedJSON(404, gin.H{"message": errss.Error()})
			return
		}
		result, err := taskcontroller.TaskUsecase.DeleteTask(Id)
		if err != nil {
			ctx.IndentedJSON(500, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(200, result)
	}
}
