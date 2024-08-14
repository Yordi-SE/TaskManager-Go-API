package usecasetest

import (
	"testing"

	"github.com/stretchr/testify/suite"
	domain "github.com/zaahidali/task_manager_api/Domain"
	usecases "github.com/zaahidali/task_manager_api/Usecases"
	"github.com/zaahidali/task_manager_api/test/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskUseCaseTestSuite defines a test suite for TaskUseCase
type TaskUseCaseTestSuite struct {
	suite.Suite
	TaskUseCase        *usecases.TaskUseCase
	mockTaskRepository *mocks.TaskRepositoryInterface
	mockUserRepository *mocks.UserRepositoryInterface
}

// SetupTest initializes the test suite
func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.mockTaskRepository = new(mocks.TaskRepositoryInterface)
	suite.mockUserRepository = new(mocks.UserRepositoryInterface)
	suite.TaskUseCase = &usecases.TaskUseCase{
		TaskRepository: suite.mockTaskRepository,
		UserRepository: suite.mockUserRepository,
	}
}

// TestGetAllTasks tests the GetAllTasks method
func (suite *TaskUseCaseTestSuite) TestGetAllTasks() {
	expectedTasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Task 1 Description"},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Task 2 Description"},
	}

	suite.mockTaskRepository.On("GetAll").Return(expectedTasks, nil)

	result, err := suite.TaskUseCase.GetAlltasks()
	suite.NoError(err)
	suite.Equal(expectedTasks, result)
	suite.mockTaskRepository.AssertExpectations(suite.T())
}

// TestGetSpecificTask tests the GetSpecificTask method
func (suite *TaskUseCaseTestSuite) TestGetSpecificTask() {
	taskID := primitive.NewObjectID()
	expectedTask := domain.Task{ID: taskID, Title: "Task 1", Description: "Task 1 Description"}

	suite.mockTaskRepository.On("GetSpecificTask", taskID).Return(expectedTask, nil)

	result, err := suite.TaskUseCase.GetSpecificTask(taskID)
	suite.NoError(err)
	suite.Equal(expectedTask, result)
	suite.mockTaskRepository.AssertExpectations(suite.T())
}

// TestCreateTask tests the CreateTask method
func (suite *TaskUseCaseTestSuite) TestCreateTask() {
	task := domain.Task{Title: "New Task", Description: "New Task Description"}
	expectedResult := primitive.NewObjectID()

	suite.mockTaskRepository.On("CreateTask", task).Return(&mongo.InsertOneResult{InsertedID: expectedResult}, nil)

	result, err := suite.TaskUseCase.CreateTask(task)
	suite.NoError(err)
	suite.Equal(expectedResult, result)
	suite.mockTaskRepository.AssertExpectations(suite.T())
}

// TestUpdateTask tests the UpdateTask method
func (suite *TaskUseCaseTestSuite) TestUpdateTask() {
	taskID := primitive.NewObjectID()
	task := domain.Task{Title: "Updated Task", Description: "Updated Task Description"}

	suite.mockTaskRepository.On("UpdateTask", taskID, task).Return(nil, nil)

	result, err := suite.TaskUseCase.UpdateTask(taskID, task)
	suite.NoError(err)
	suite.Nil(result)
	suite.mockTaskRepository.AssertExpectations(suite.T())
}

// TestDeleteTask tests the DeleteTask method
func (suite *TaskUseCaseTestSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID()
	expectedResult := int64(1)

	suite.mockTaskRepository.On("DeleteTask", taskID).Return(&mongo.DeleteResult{DeletedCount: expectedResult}, nil)

	result, err := suite.TaskUseCase.DeleteTask(taskID)
	suite.NoError(err)
	suite.Equal(expectedResult, result)
	suite.mockTaskRepository.AssertExpectations(suite.T())
}

// TestDeleteTaskNotFound tests the DeleteTask method when the task is not found
func (suite *TaskUseCaseTestSuite) TestDeleteTaskNotFound() {
	taskID := primitive.NewObjectID()

	suite.mockTaskRepository.On("DeleteTask", taskID).Return(&mongo.DeleteResult{DeletedCount: 0}, nil)

	result, err := suite.TaskUseCase.DeleteTask(taskID)
	suite.Error(err)
	suite.Nil(result)
	suite.EqualError(err, "task not found")
	suite.mockTaskRepository.AssertExpectations(suite.T())
}

// Run the test suite
func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}
