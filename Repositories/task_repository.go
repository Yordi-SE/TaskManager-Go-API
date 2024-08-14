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

// Abstracts the data access logic.
// The repository pattern is a design pattern that isolates the data access logic from the rest of the application.
type TaskRepository struct {
	TaskCollection *mongo.Collection
	UserCollection *mongo.Collection
	Database       *mongo.Database
}

type TaskRepositoryInterface interface {
	GetAll() ([]domain.Task, error)
	GetSpecificTask(id primitive.ObjectID) (domain.Task, error)
	CreateTask(task domain.Task) (*mongo.InsertOneResult, error)
	UpdateTask(id primitive.ObjectID, task domain.Task) (*mongo.UpdateResult, error)
	DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Count(col *mongo.Collection) (int64, error)
}

func (taskrepository *TaskRepository) GetAll() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var results []domain.Task

	findOptions := options.Find()
	cur, err := taskrepository.TaskCollection.Find(ctx, bson.M{}, findOptions)
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

func (taskrepository *TaskRepository) GetSpecificTask(id primitive.ObjectID) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var task domain.Task
	err := taskrepository.TaskCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

// create task
func (taskrepository *TaskRepository) CreateTask(task domain.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := taskrepository.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// update task
func (taskrepository *TaskRepository) UpdateTask(id primitive.ObjectID, task domain.Task) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := taskrepository.TaskCollection.UpdateByID(ctx, id, bson.D{
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
func (taskrepository *TaskRepository) DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := taskrepository.TaskCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// count specific collection
func (taskrepository *TaskRepository) Count(col *mongo.Collection) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := col.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
