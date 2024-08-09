package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
}

var Tasks = []Task{
	{Title: "Task 1", Description: "Description 1"},
	{Title: "Task 2", Description: "Description 2"},
	{Title: "Task 3", Description: "Description 3"},
}
