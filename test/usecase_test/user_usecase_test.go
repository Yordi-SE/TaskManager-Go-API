package usecasetest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	domain "github.com/zaahidali/task_manager_api/Domain"
	usecases "github.com/zaahidali/task_manager_api/Usecases"
	"github.com/zaahidali/task_manager_api/test/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	UserUseCase        *usecases.UserUseCase
	MockTaskRepository *mocks.TaskRepositoryInterface
	MockUserRepository *mocks.UserRepositoryInterface
	MockInfrastructure *mocks.InfrastructureInterface
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.MockTaskRepository = new(mocks.TaskRepositoryInterface)
	suite.MockUserRepository = new(mocks.UserRepositoryInterface)
	suite.MockInfrastructure = new(mocks.InfrastructureInterface)

	suite.UserUseCase = &usecases.UserUseCase{
		TaskRepository: suite.MockTaskRepository,
		UserRepository: suite.MockUserRepository,
		Infrastracture: suite.MockInfrastructure,
	}
}

// Test Register method
func (suite *UserUseCaseTestSuite) TestRegister_Success() {
	user := domain.User{
		UserName: "testuser",
		Password: "hashedpassword",
		Role:     "admin",
	}
	var col *mongo.Collection = nil
	hashedPassword := "hashedpassword"
	suite.MockInfrastructure.On("HashPassword", user.Password).Return(hashedPassword, nil)
	suite.MockTaskRepository.On("Count", col).Return(int64(0), nil)
	suite.MockUserRepository.On("FindUserByName", user.UserName).Return(domain.User{}, mongo.ErrNoDocuments)
	suite.MockUserRepository.On("CreateUser", user).Return(&mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil)

	insertedID, err := suite.UserUseCase.Register(user)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), insertedID)
}

func (suite *UserUseCaseTestSuite) TestRegister_UserAlreadyExists() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	existingUser := domain.User{
		UserName: "testuser",
	}
	hashedPassword := "hashedpassword"
	var col *mongo.Collection = nil

	suite.MockInfrastructure.On("HashPassword", user.Password).Return(hashedPassword, nil)
	suite.MockTaskRepository.On("Count", col).Return(int64(0), nil)

	suite.MockUserRepository.On("FindUserByName", user.UserName).Return(existingUser, nil)

	insertedID, err := suite.UserUseCase.Register(user)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "user already exists")
	assert.Nil(suite.T(), insertedID)
}

func (suite *UserUseCaseTestSuite) TestPromote_UserNotFound() {
	userID := primitive.NewObjectID()

	suite.MockUserRepository.On("Promote", userID).Return(&mongo.UpdateResult{MatchedCount: 0}, nil)

	err := suite.UserUseCase.Promote(userID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "user not found")
}

func (suite *UserUseCaseTestSuite) TestPromote_UserIsAdmin() {
	userID := primitive.NewObjectID()

	suite.MockUserRepository.On("Promote", userID).Return(&mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 0}, nil)

	err := suite.UserUseCase.Promote(userID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "user is admin")
}

// Test Login method
func (suite *UserUseCaseTestSuite) TestLogin_Success() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	foundUser := domain.User{
		UserName: "testuser",
		Password: "hashedpassword",
		Role:     "user",
	}

	suite.MockUserRepository.On("FindUserByName", user.UserName).Return(foundUser, nil)
	suite.MockInfrastructure.On("ComparePasswords", foundUser.Password, user.Password).Return(nil)
	suite.MockInfrastructure.On("GenerateToken", foundUser.UserName, foundUser.ID, foundUser.Role).Return("jwtToken", nil)

	token, err := suite.UserUseCase.Login(user)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "jwtToken", token)
}

func (suite *UserUseCaseTestSuite) TestLogin_InvalidPassword() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	foundUser := domain.User{
		UserName: "testuser",
		Password: "hashedpassword",
		Role:     "user",
	}

	suite.MockUserRepository.On("FindUserByName", user.UserName).Return(foundUser, nil)
	suite.MockInfrastructure.On("ComparePasswords", foundUser.Password, user.Password).Return(errors.New("invalid password"))

	token, err := suite.UserUseCase.Login(user)

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), token)
}

func (suite *UserUseCaseTestSuite) TestLogin_UserNotFound() {
	user := domain.User{
		UserName: "testuser",
		Password: "testpassword",
	}

	suite.MockUserRepository.On("FindUserByName", user.UserName).Return(domain.User{}, mongo.ErrNoDocuments)

	token, err := suite.UserUseCase.Login(user)

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), token)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
