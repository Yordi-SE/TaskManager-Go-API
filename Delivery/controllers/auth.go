package controllers

import (
	"fmt"
	"net/http"

	domain "github.com/zaahidali/task_manager_api/Domain"
	usecases "github.com/zaahidali/task_manager_api/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandler struct {
	UserUsecase usecases.UserUseCaseInterface
	TaskUsecase usecases.TaskUseCaseInterface
}

// auth handler

func (authhandler *AuthHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user domain.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		}
		result, errs := authhandler.UserUsecase.Register(user)
		if errs != nil {
			ctx.JSON(500, gin.H{"error": errs.Error()})
			return
		}
		ctx.JSON(200, result)

	}
}

// login handler

func (authhandler *AuthHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user domain.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		}
		jwtToken, errs := authhandler.UserUsecase.Login(user)
		if errs != nil {
			ctx.IndentedJSON(http.StatusForbidden, errs.Error())
			return
		}
		ctx.IndentedJSON(200, gin.H{"message": "Login successful", "token": jwtToken})

	}

}

// promote user
func (authhandler *AuthHandler) Promote() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user struct {
			UserId string `json:"user_id"`
		}
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		}
		fmt.Println(user.UserId)
		Id, errss := primitive.ObjectIDFromHex(user.UserId)
		if errss != nil {
			ctx.IndentedJSON(400, gin.H{"message": errss.Error()})
			return
		}
		err := authhandler.UserUsecase.Promote(Id)
		if err != nil {
			ctx.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(200, gin.H{"message": "promotion was successful"})

	}
}
