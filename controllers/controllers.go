package controllers

import (
	"strconv"

	"github.com/zaahidali/task_manager_api/models"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/data"
)

func GetTasks(ctx *gin.Context) {
	ctx.IndentedJSON(200, data.Tasks)
}

func GetTasksId(ctx *gin.Context) {
	id, _ := (strconv.Atoi(ctx.Param("id")))

	for _, task := range data.Tasks {
		if task.ID == uint(id) {
			ctx.IndentedJSON(200, task)
			return
		}
	}
	ctx.IndentedJSON(404, gin.H{"message": "Task not found"})

}

func CreateTask(ctx *gin.Context) {
	var task models.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": "Invalid request"})
		return
	}
	task.ID = uint(len(data.Tasks) + 1)
	data.Tasks = append(data.Tasks, task)
	ctx.IndentedJSON(201, task)
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": "Invalid request"})
		return
	}
	for i, task := range data.Tasks {
		if task.ID == uint(idInt) {
			var updatedTask models.Task
			err := ctx.BindJSON(&updatedTask)
			if err != nil {
				ctx.IndentedJSON(400, gin.H{"message": "Invalid request"})
				return
			}
			updatedTask.ID = task.ID
			data.Tasks[i] = updatedTask
			ctx.IndentedJSON(200, updatedTask)
			return
		}
	}
	ctx.IndentedJSON(404, gin.H{"message": "Task not found"})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": "Invalid request"})
		return
	}
	for i, task := range data.Tasks {
		if task.ID == uint(idInt) {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			ctx.IndentedJSON(200, gin.H{"message": "Task deleted"})
			return
		}
	}
	ctx.IndentedJSON(404, gin.H{"message": "Task not found"})
}
