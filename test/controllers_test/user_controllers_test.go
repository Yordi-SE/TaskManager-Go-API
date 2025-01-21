package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zaahidali/task_manager_api/Delivery/controllers"
	domain "github.com/zaahidali/task_manager_api/Domain"
	"github.com/zaahidali/task_manager_api/test/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerTestSuite struct {
	suite.Suite
	Controller               *controllers.TaskHandler
	UserController           *controllers.AuthHandler
	mockUserUseCaseInterface *mocks.UserUseCaseInterface
	mockTaskUseCaseInterface *mocks.TaskUseCaseInterface
	router                   *gin.Engine
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUserUseCaseInterface = new(mocks.UserUseCaseInterface)
	suite.mockTaskUseCaseInterface = new(mocks.TaskUseCaseInterface)
	suite.UserController = &controllers.AuthHandler{
		UserUsecase: suite.mockUserUseCaseInterface,
		TaskUsecase: suite.mockTaskUseCaseInterface,
	}
	suite.router = gin.Default()
	suite.router.POST("/register", suite.UserController.Register())
	suite.router.POST("/login", suite.UserController.Login())
	suite.router.POST("/promote", suite.UserController.Promote())
}

// Register Test Cases
func (suite *UserControllerTestSuite) TestRegister_Valid() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}
	userId := primitive.NewObjectID()

	suite.mockUserUseCaseInterface.On("Register", user).Return(userId, nil)

	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockUserUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestRegister_InvalidJSON() {
	invalidJSON := `{"UserName":"testuser", "Password":}`

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockUserUseCaseInterface.AssertNotCalled(suite.T(), "Register", mock.Anything)
}

func (suite *UserControllerTestSuite) TestRegister_Error() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	suite.mockUserUseCaseInterface.On("Register", user).Return(domain.User{}, errors.New("registration error"))

	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
	suite.Contains(w.Body.String(), "error")
	suite.mockUserUseCaseInterface.AssertExpectations(suite.T())
}

// Login Test Cases
func (suite *UserControllerTestSuite) TestLogin_Valid() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	suite.mockUserUseCaseInterface.On("Login", user).Return("mockToken", nil)

	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "token")
	suite.mockUserUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogin_InvalidJSON() {
	invalidJSON := `{"UserName":"testuser", "Password":}`

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockUserUseCaseInterface.AssertNotCalled(suite.T(), "Login", mock.Anything)
}

func (suite *UserControllerTestSuite) TestLogin_Error() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	suite.mockUserUseCaseInterface.On("Login", user).Return("", errors.New("login error"))

	userJSON, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusForbidden, w.Code)
	suite.Contains(w.Body.String(), "login error")
	suite.mockUserUseCaseInterface.AssertExpectations(suite.T())
}

// Promote Test Cases
func (suite *UserControllerTestSuite) TestPromote_Valid() {
	userID := primitive.NewObjectID().Hex()
	promoteRequest := struct {
		UserId string `json:"user_id"`
	}{
		UserId: userID,
	}

	suite.mockUserUseCaseInterface.On("Promote", mock.AnythingOfType("primitive.ObjectID")).Return(nil)

	reqBody, _ := json.Marshal(promoteRequest)
	req, _ := http.NewRequest(http.MethodPost, "/promote", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.Contains(w.Body.String(), "promotion was successful")
	suite.mockUserUseCaseInterface.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestPromote_InvalidJSON() {
	invalidJSON := `{"user_id":}`

	req, _ := http.NewRequest(http.MethodPost, "/promote", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockUserUseCaseInterface.AssertNotCalled(suite.T(), "Promote", mock.Anything)
}

func (suite *UserControllerTestSuite) TestPromote_InvalidID() {
	invalidID := "invalid_id"
	promoteRequest := struct {
		UserId string `json:"user_id"`
	}{
		UserId: invalidID,
	}

	reqBody, _ := json.Marshal(promoteRequest)
	req, _ := http.NewRequest(http.MethodPost, "/promote", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "message")
	suite.mockUserUseCaseInterface.AssertNotCalled(suite.T(), "Promote", mock.Anything)
}

func (suite *UserControllerTestSuite) TestPromote_Error() {
	userID := primitive.NewObjectID()
	promoteRequest := struct {
		UserId string `json:"user_id"`
	}{
		UserId: userID.Hex(),
	}

	suite.mockUserUseCaseInterface.On("Promote", userID).Return(errors.New("promotion error"))

	reqBody, _ := json.Marshal(promoteRequest)
	req, _ := http.NewRequest(http.MethodPost, "/promote", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
	suite.Contains(w.Body.String(), "promotion error")
	suite.mockUserUseCaseInterface.AssertExpectations(suite.T())
}

// Run the test suite
func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
