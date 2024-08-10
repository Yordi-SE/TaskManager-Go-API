package repositories

import (
	"context"
	"time"

	domain "github.com/zaahidali/task_manager_api/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Abstracts the data access logic.
//The repository pattern is a design pattern that isolates the data access logic from the rest of the application.

type TaskRepositoryInterface interface {
	GetAll() ([]domain.Task, error)
	GetSpecificTask(id primitive.ObjectID) (domain.Task, error)
	CreateTask(task domain.Task) (*mongo.InsertOneResult, error)
	UpdateTask(id primitive.ObjectID, task domain.Task) (*mongo.UpdateResult, error)
	DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Count(col *mongo.Collection) (int64, error)
}

var TaskRepository TaskRepositoryInterface

func GetAll() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var results []domain.Task

	findOptions := options.Find()
	cur, err := Collections.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var elem domain.Task
		err := cur.Decode(&elem)
		if err != nil {
			return []domain.Task{}, err
		}

		results = append(results, elem)
	}
	return results, nil
}

func GetSpecificTask(id primitive.ObjectID) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var task domain.Task
	err := Collections.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

// create task
func CreateTask(task domain.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := Collections.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// update task
func UpdateTask(id primitive.ObjectID, task domain.Task) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := Collections.UpdateByID(ctx, id, bson.D{
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
func DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := Collections.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// count specific collection
func Count(col *mongo.Collection) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := col.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
