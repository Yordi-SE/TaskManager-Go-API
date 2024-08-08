package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zaahidali/task_manager_api/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// setup connection
var Collections *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {

		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Collections = client.Database("task_manager").Collection("tasks")
	tasks := []interface{}{
		data.Tasks[0],
		data.Tasks[1],
		data.Tasks[2],
	}
	fmt.Println("Collection instance created!")
	Collections.DeleteMany(context.TODO(), bson.D{{}})
	Collections.InsertMany(context.TODO(), tasks)
}

// Task struct
