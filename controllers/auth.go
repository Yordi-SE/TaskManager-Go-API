package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// auth handler

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	result, errs := models.UserCollection.InsertOne(ctx, user)
	if errs != nil {
		ctx.JSON(500, gin.H{"error": errs.Error()})
		return
	}
	ctx.JSON(200, result.InsertedID)

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
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		ctx.IndentedJSON(401, gin.H{"message": "Invalid credentials"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": result.ID,
		"email":   result.Email,
	})
	jwtToken, err := token.SignedString([]byte(secret))
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(200, gin.H{"message": "Login successful", "token": jwtToken})

}
