package controllers

import (
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/data"
)

func GetTasks(ctx *gin.Context) {
	var results []data.Task

	findOptions := options.Find()
	cur, err := models.Collections.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	for cur.Next(ctx) {
		var elem data.Task
		err := cur.Decode(&elem)
		if err != nil {
			ctx.IndentedJSON(500, gin.H{"message": err.Error()})
			return
		}

		results = append(results, elem)
	}
	ctx.IndentedJSON(200, results)
}

func GetTasksId(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, errs := primitive.ObjectIDFromHex(id)
	if errs != nil {
		ctx.IndentedJSON(404, gin.H{"message": errs.Error()})
	}
	var tasks data.Task
	err := models.Collections.FindOne(ctx, bson.D{{Key: "_id", Value: Id}}).Decode(&tasks)
	if err != nil {
		ctx.IndentedJSON(404, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, tasks)

}

func CreateTask(ctx *gin.Context) {
	var task data.Task
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	result, err := models.Collections.InsertOne(ctx, task)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(201, result.InsertedID)
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, errss := primitive.ObjectIDFromHex(id)
	if errss != nil {
		ctx.IndentedJSON(404, gin.H{"message": errss.Error()})
		return
	}
	var task data.Task
	errs := ctx.BindJSON(&task)
	if errs != nil {
		ctx.IndentedJSON(400, gin.H{"message": errs.Error()})
		return
	}
	result, err := models.Collections.UpdateByID(ctx, Id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "title", Value: task.Title},
				{Key: "description", Value: task.Description},
			},
		},
	})
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
	result, err := models.Collections.DeleteOne(ctx, bson.M{"_id": Id})
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, result.DeletedCount)
}
