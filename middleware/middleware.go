package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		errss := godotenv.Load()
		jwtSecret := os.Getenv("jwtSecret")
		fmt.Println("jwtSecret", jwtSecret)
		if errss != nil {
			ctx.JSON(500, gin.H{
				"message": "server error"})
			ctx.Abort()
			return
		}
		// check if user is authenticated
		// if not, return 401
		// else, continue
		tokens := ctx.GetHeader("Authorization")
		if tokens == "" {
			ctx.JSON(400, gin.H{"message": "Invalid Token!"})
			ctx.Abort()
			return
		}
		authtokens := strings.Split(tokens, " ")
		fmt.Println(authtokens)
		if len(authtokens) != 2 || strings.ToLower(authtokens[0]) != "bearer" {
			ctx.JSON(401, gin.H{"error": "Invalid authorization header"})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(authtokens[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])

			}
			return []byte(jwtSecret), nil
		})
		fmt.Println(err)
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"error": "Invalid JWT", "valid": token.Valid})
			ctx.Abort()
			return
		}
		ctx.Next()

	}
}
