package Infrastructure

// Functions to generate and validate JWT tokens.
import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GenerateToken generates a JWT token
func (*Infrastructure) GenerateToken(user_name string, id primitive.ObjectID, role string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	secret := os.Getenv("jwtSecret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name": user_name,
		"user_id":   id,
		"role":      role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

// ValidateToken validates a JWT token
func (*Infrastructure) ValidateToken(tokenString string) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	secret := os.Getenv("jwtSecret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}
