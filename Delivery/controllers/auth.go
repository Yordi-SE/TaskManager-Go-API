package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/zaahidali/task_manager_api/Infrastructure"
	usecases "github.com/zaahidali/task_manager_api/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// auth handler

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	result, errs := usecases.Register(ctx, user)
	if errs != nil {
		ctx.JSON(500, gin.H{"error": errs.Error()})
		return
	}
	ctx.JSON(200, result)

}

// login handler

func Login(ctx *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("jwtSecret")
	fmt.Println("jwtSecret", secret)
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	var result models.User
	errs := models.UserCollection.FindOne(ctx, bson.D{{Key: "email", Value: user.Email}}).Decode(&result)
	if errs != nil {
		ctx.IndentedJSON(404, gin.H{"message": err.Error()})
		return
	}
	err = Infrastructure.ComparePasswords(result.Password, user.Password)
	if err != nil {
		ctx.IndentedJSON(401, gin.H{"message": "Invalid credentials"})
		return
	}
	jwtToken, err := Infrastructure.GenerateToken(result.Email, result.ID, result.Role)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, gin.H{"message": "Login successful", "token": jwtToken})

}

// promote user
func Promote(ctx *gin.Context) {
	var user struct {
		user_id string
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	Id, errss := primitive.ObjectIDFromHex(user.user_id)
	if errss != nil {
		ctx.IndentedJSON(400, gin.H{"message": errss.Error()})
		return
	}
	err := usecases.Promote(ctx, Id)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(400, gin.H{"message": "promotion was successful"})

}
