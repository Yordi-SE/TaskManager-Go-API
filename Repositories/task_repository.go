package repositories

import (
	"github.com/gin-gonic/gin"
	domain "github.com/zaahidali/task_manager_api/Domain"
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Abstracts the data access logic.
//The repository pattern is a design pattern that isolates the data access logic from the rest of the application.

func GetAll(ctx *gin.Context) ([]domain.Task, error) {

	var results []domain.Task

	findOptions := options.Find()
	cur, err := models.Collections.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem domain.Task
		err := cur.Decode(&elem)
		if err != nil {
			return []domain.Task{}, err
		}

		results = append(results, elem)
	}
	return results, nil
}

func GetSpecificTask(ctx *gin.Context, id primitive.ObjectID) (domain.Task, error) {
	var task domain.Task
	err := models.Collections.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

// create task
func CreateTask(ctx *gin.Context, task domain.Task) (*mongo.InsertOneResult, error) {
	result, err := models.Collections.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// update task
func UpdateTask(ctx *gin.Context, id primitive.ObjectID, task domain.Task) (*mongo.UpdateResult, error) {
	result, err := models.Collections.UpdateByID(ctx, id, bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "title", Value: task.Title},
				{Key: "description", Value: task.Description},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// delete task
func DeleteTask(ctx *gin.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := models.Collections.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// count specific collection
func Count(ctx *gin.Context, col *mongo.Collection) (int64, error) {
	count, err := col.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
