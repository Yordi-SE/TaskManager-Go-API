package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zaahidali/task_manager_api/Delivery/controllers"
	domain "github.com/zaahidali/task_manager_api/Domain"
	"github.com/zaahidali/task_manager_api/test/mocks"
)

// TaskControllerTestSuite defines the test suite for TaskController
type TaskControllerTestSuite struct {
	suite.Suite
	Controller               *controllers.TaskHandler
	UserController           *controllers.AuthHandler
	mockTaskUseCaseInterface *mocks.TaskUseCaseInterface
	mockUserUseCaseInterface *mocks.UserUseCaseInterface
	router                   *gin.Engine
	testingServer            *httptest.Server
}

func (suite *TaskControllerTestSuite) SetupTest() {
	gin.SetMode("test")
	suite.mockUserUseCaseInterface = new(mocks.UserUseCaseInterface)
	suite.mockTaskUseCaseInterface = new(mocks.TaskUseCaseInterface)
	suite.Controller = &controllers.TaskHandler{
		UserUsecase: suite.mockUserUseCaseInterface,
		TaskUsecase: suite.mockTaskUseCaseInterface,
	}
	suite.router = gin.Default()

	// Set up routes
	suite.router.GET("/tasks", suite.Controller.GetTasks())
	suite.router.GET("/tasks/:id", suite.Controller.GetTasksId())
	suite.router.POST("/tasks", suite.Controller.CreateTask())
	suite.router.PUT("/tasks/:id", suite.Controller.UpdateTask())
	suite.router.DELETE("/tasks/:id", suite.Controller.DeleteTask())
	testingServer := httptest.NewServer(suite.router)
	suite.testingServer = testingServer

}

func (suite *TaskControllerTestSuite) TestGetTasks_NoTasks() {
	suite.mockTaskUseCaseInterface.On("GetAlltasks").Return([]domain.Task{}, nil)

	res, err := http.Get(fmt.Sprintf("%s/tasks", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")
	defer res.Body.Close()

	response := []domain.Task{}
	json.NewDecoder(res.Body).Decode(&response)

	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(response, []domain.Task{})

	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetTasks_WithTasks() {
	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Description 2"},
	}
	suite.mockTaskUseCaseInterface.On("GetAlltasks").Return(tasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	expectedResponse, _ := json.Marshal(tasks)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), string(expectedResponse), w.Body.String())
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetTasks_Error() {
	suite.mockTaskUseCaseInterface.On("GetAlltasks").Return(nil, assert.AnError)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	assert.Contains(suite.T(), w.Body.String(), assert.AnError.Error())

	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetTasksId_Success() {
	// Arrange
	taskID := primitive.NewObjectID()
	expectedTask := domain.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "description",
	}

	suite.mockTaskUseCaseInterface.On("GetSpecificTask", taskID).Return(expectedTask, nil)

	// Act
	req, _ := http.NewRequest(http.MethodGet, "/tasks/"+taskID.Hex(), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	suite.Equal(http.StatusOK, w.Code)
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetTasksId_InvalidID() {
	// Arrange
	invalidID := "invalidID"
	suite.mockTaskUseCaseInterface.On("GetSpecificTask", invalidID).Return(nil, fmt.Errorf("not found"))

	// Act
	req, _ := http.NewRequest("GET", "/tasks/"+invalidID, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *TaskControllerTestSuite) TestGetTasksId_TaskNotFound() {
	// Arrange
	taskID := primitive.NewObjectID()

	suite.mockTaskUseCaseInterface.On("GetSpecificTask", taskID).Return(domain.Task{}, fmt.Errorf("task not found"))

	// Act
	req, _ := http.NewRequest("GET", "/tasks/"+taskID.Hex(), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Assert
	suite.Equal(http.StatusNotFound, w.Code)
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestCreateTask_Valid() {
	// Define input and expected output
	task := domain.Task{
		Title:       "Test Task",
		Description: "Test Task Description",
	}
	taskID := primitive.NewObjectID()

	suite.mockTaskUseCaseInterface.On("CreateTask", mock.AnythingOfType("domain.Task")).Return(taskID, nil)

	// Convert task to JSON
	taskJSON, _ := json.Marshal(task)

	// Create request and record response
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	suite.router.ServeHTTP(w, req)
	fmt.Println(w.Body)
	// Assertions
	suite.Equal(http.StatusCreated, w.Code)
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

// TestCreateTask_InvalidJSON tests invalid JSON input
func (suite *TaskControllerTestSuite) TestCreateTask_InvalidJSON() {
	invalidJSON := `{"Title":"Test Task", "Description":}` // Invalid JSON

	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	suite.router.ServeHTTP(w, req)

	// Assertions
	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockTaskUseCaseInterface.AssertNotCalled(suite.T(), "CreateTask", mock.Anything)
}

// TestCreateTask_Error tests when the use case returns an error
func (suite *TaskControllerTestSuite) TestCreateTask_Error() {
	task := domain.Task{
		Title:       "Test Task",
		Description: "Test Task Description",
	}

	// Mocking the CreateTask function to return an error
	suite.mockTaskUseCaseInterface.On("CreateTask", mock.AnythingOfType("domain.Task")).Return(nil, errors.New("database error"))

	// Convert task to JSON
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	suite.router.ServeHTTP(w, req)

	// Assertions
	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestUpdateTask_Valid() {
	task := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mockID := primitive.NewObjectID()
	suite.mockTaskUseCaseInterface.On("UpdateTask", mockID, task).Return(task, nil)

	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+mockID.Hex(), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestUpdateTask_InvalidID() {
	task := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/invalid_id", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockTaskUseCaseInterface.AssertNotCalled(suite.T(), "UpdateTask", mock.Anything)
}

func (suite *TaskControllerTestSuite) TestUpdateTask_InvalidJSON() {
	mockID := primitive.NewObjectID()
	invalidJSON := `{"Title":"Updated Task", "Description":}`

	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+mockID.Hex(), bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockTaskUseCaseInterface.AssertNotCalled(suite.T(), "UpdateTask", mock.Anything)
}

func (suite *TaskControllerTestSuite) TestUpdateTask_Error() {
	task := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mockID := primitive.NewObjectID()
	suite.mockTaskUseCaseInterface.On("UpdateTask", mockID, task).Return(domain.Task{}, errors.New("database error"))

	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+mockID.Hex(), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

// DeleteTask Test Cases
func (suite *TaskControllerTestSuite) TestDeleteTask_Valid() {
	mockID := primitive.NewObjectID()
	suite.mockTaskUseCaseInterface.On("DeleteTask", mockID).Return(domain.Task{}, nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+mockID.Hex(), nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestDeleteTask_InvalidID() {
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/invalid_id", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockTaskUseCaseInterface.AssertNotCalled(suite.T(), "DeleteTask", mock.Anything)
}

func (suite *TaskControllerTestSuite) TestDeleteTask_Error() {
	mockID := primitive.NewObjectID()
	suite.mockTaskUseCaseInterface.On("DeleteTask", mockID).Return(domain.Task{}, errors.New("database error"))

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+mockID.Hex(), nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockTaskUseCaseInterface.AssertExpectations(suite.T())
}
func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
