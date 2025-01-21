package Infrastructure

import "go.mongodb.org/mongo-driver/bson/primitive"

type Infrastructure struct {
}
type InfrastructureInterface interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword string, password string) error
	GenerateToken(user_name string, id primitive.ObjectID, role string) (string, error)
	ValidateToken(tokenString string) (bool, error)
}
