package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserName string             `json:"user_name" bson:"username" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
	Role     string             `json:"role" bson:"role"`
}

// setup connection

// Task struct
