package repositorytest

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	domain "github.com/zaahidali/task_manager_api/Domain"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	client         *mongo.Client
	db             *mongo.Database
	taskCollection *mongo.Collection
	repo           repositories.TaskRepository
}

// SetupSuite runs before all tests
func (suite *TaskRepositoryTestSuite) SetupTest() {

	var db_url = os.Getenv("DB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(db_url).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Connect to MongoDB
	suite.client = client
	suite.NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	suite.NoError(err)

	// Ping the database to verify the connection
	err = suite.client.Ping(ctx, readpref.Primary())
	suite.NoError(err)

	// Use a specific test database
	Database := client.Database("task_manager_test")

	suite.db = Database
	suite.taskCollection = suite.db.Collection("tasks")
	usercollection := suite.db.Collection("users")

	// Initialize the repository
	suite.repo = repositories.TaskRepository{
		TaskCollection: suite.taskCollection,
		UserCollection: usercollection,
		Database:       suite.db,
	}
}

// TearDownSuite runs after all tests
func (suite *TaskRepositoryTestSuite) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Drop the database to clean up after tests
	err := suite.db.Drop(ctx)
	suite.NoError(err)

	// Disconnect the MongoDB client
	err = suite.client.Disconnect(ctx)
	suite.NoError(err)
}

// TestGetAll tests the GetAll method
func (suite *TaskRepositoryTestSuite) TestGetAll() {
	// Insert some sample data
	_, err := suite.taskCollection.InsertOne(context.Background(), domain.Task{
		Title:       "Test Task 1",
		Description: "Description 1",
	})
	suite.NoError(err)

	_, err = suite.taskCollection.InsertOne(context.Background(), domain.Task{
		Title:       "Test Task 2",
		Description: "Description 2",
	})
	suite.NoError(err)

	// Test GetAll method
	tasks, err := suite.repo.GetAll()
	suite.NoError(err)
	suite.Equal(2, len(tasks))
}

// TestGetSpecificTask tests the GetSpecificTask method
func (suite *TaskRepositoryTestSuite) TestGetSpecificTask() {
	task := domain.Task{
		Title:       "Specific Task",
		Description: "Specific Description",
	}
	insertResult, err := suite.taskCollection.InsertOne(context.Background(), task)
	suite.NoError(err)

	id := insertResult.InsertedID.(primitive.ObjectID)

	// Test GetSpecificTask method
	retrievedTask, err := suite.repo.GetSpecificTask(id)
	suite.NoError(err)
	suite.Equal("Specific Task", retrievedTask.Title)
}

// TestCreateTask tests the CreateTask method
func (suite *TaskRepositoryTestSuite) TestCreateTask() {
	task := domain.Task{
		Title:       "New Task",
		Description: "New Task Description",
	}

	// Test CreateTask method
	insertResult, err := suite.repo.CreateTask(task)
	suite.NoError(err)
	suite.NotNil(insertResult.InsertedID)
}

// TestUpdateTask tests the UpdateTask method
func (suite *TaskRepositoryTestSuite) TestUpdateTask() {
	task := domain.Task{
		Title:       "Task to Update",
		Description: "Original Description",
	}
	insertResult, err := suite.taskCollection.InsertOne(context.Background(), task)
	suite.NoError(err)

	id := insertResult.InsertedID.(primitive.ObjectID)

	updatedTask := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	// Test UpdateTask method
	updateResult, err := suite.repo.UpdateTask(id, updatedTask)
	suite.NoError(err)
	suite.Equal(int64(1), updateResult.ModifiedCount)

	// Verify the update
	retrievedTask, err := suite.repo.GetSpecificTask(id)
	suite.NoError(err)
	suite.Equal("Updated Task", retrievedTask.Title)
}

// TestDeleteTask tests the DeleteTask method
func (suite *TaskRepositoryTestSuite) TestDeleteTask() {
	task := domain.Task{
		Title:       "Task to Delete",
		Description: "Delete Description",
	}
	insertResult, err := suite.taskCollection.InsertOne(context.Background(), task)
	suite.NoError(err)

	id := insertResult.InsertedID.(primitive.ObjectID)

	// Test DeleteTask method
	deleteResult, err := suite.repo.DeleteTask(id)
	suite.NoError(err)
	suite.Equal(int64(1), deleteResult.DeletedCount)

	// Verify the deletion
	_, err = suite.repo.GetSpecificTask(id)
	suite.Error(err)
}

// TestCount tests the Count method
func (suite *TaskRepositoryTestSuite) TestCount() {
	// Insert some sample data
	_, err := suite.taskCollection.InsertOne(context.Background(), domain.Task{
		Title:       "Task 1",
		Description: "Description 1",
	})
	suite.NoError(err)

	_, err = suite.taskCollection.InsertOne(context.Background(), domain.Task{
		Title:       "Task 2",
		Description: "Description 2",
	})
	suite.NoError(err)

	// Test Count method
	count, err := suite.repo.Count(suite.taskCollection)
	suite.NoError(err)
	suite.Equal(int64(2), count)
}

// Run the test suite
func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
