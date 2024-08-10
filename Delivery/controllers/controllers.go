package controllers

import (
	usecases "github.com/zaahidali/task_manager_api/Usecases"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	domain "github.com/zaahidali/task_manager_api/Domain"
)

func GetTasks(ctx *gin.Context) {
	result, err := usecases.GetAlltasks()
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, result)
}

func GetTasksId(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, errs := primitive.ObjectIDFromHex(id)
	if errs != nil {
		ctx.IndentedJSON(404, gin.H{"message": errs.Error()})
	}
	tasks, err := usecases.GetSpecificTask(Id)

	if err != nil {
		ctx.IndentedJSON(404, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, tasks)

}

func CreateTask(ctx *gin.Context) {
	var task domain.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	result, err := usecases.CreateTask(task)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(201, result)
}

func UpdateTask(ctx *gin.Context) {
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
	result, err := usecases.UpdateTask(Id, task)

	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(200, result)
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, errss := primitive.ObjectIDFromHex(id)
	if errss != nil {
		ctx.IndentedJSON(404, gin.H{"message": errss.Error()})
		return
	}
	result, err := usecases.DeleteTask(Id)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, result)
}
