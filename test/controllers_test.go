package testControllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zaahidali/task_manager_api/Delivery/controllers"
	domain "github.com/zaahidali/task_manager_api/Domain"
	"github.com/zaahidali/task_manager_api/test/mocks"
)

// TaskControllerTestSuite defines the test suite for TaskController
type TaskControllerTestSuite struct {
	suite.Suite
	mockTaskUseCaseInterface *mocks.TaskUseCaseInterface
	router                   *gin.Engine
}

func (suite *TaskControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockTaskUseCaseInterface = new(mocks.TaskUseCaseInterface)
	suite.router = gin.Default()

	// Set up routes
	suite.router.GET("/tasks", controllers.GetTasks)
	suite.router.GET("/tasks/:id", controllers.GetTasksId)
	suite.router.POST("/tasks", controllers.CreateTask)
	suite.router.PUT("/tasks/:id", controllers.UpdateTask)
	suite.router.DELETE("/tasks/:id", controllers.DeleteTask)
}

func (suite *TaskControllerTestSuite) TestGetTasks_NoTasks() {
	suite.mockTaskUseCaseInterface.On("GetAlltasks").Return([]domain.Task{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `[]`, w.Body.String())
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

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
