package controllers

import (
	"fmt"
	"net/http"

	domain "github.com/zaahidali/task_manager_api/Domain"
	usecases "github.com/zaahidali/task_manager_api/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// auth handler

func Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	result, errs := usecases.Register(user)
	if errs != nil {
		ctx.JSON(500, gin.H{"error": errs.Error()})
		return
	}
	ctx.JSON(200, result)

}

// login handler

func Login(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(user.Password)

	jwtToken, errs := usecases.Login(user)
	if errs != nil {
		ctx.IndentedJSON(http.StatusForbidden, errs.Error())
		return
	}
	ctx.IndentedJSON(200, gin.H{"message": "Login successful", "token": jwtToken})

}

// promote user
func Promote(ctx *gin.Context) {
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
	err := usecases.Promote(Id)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, gin.H{"message": "promotion was successful"})

}
