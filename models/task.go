package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	domain "github.com/zaahidali/task_manager_api/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// setup connection
var Collections *mongo.Collection
var Database *mongo.Database
var UserCollection *mongo.Collection

func init() {
	Database = ConnectDB()
	Collections = Database.Collection("tasks")
	UserCollection = Database.Collection("User")
	tasks := []interface{}{
		domain.Tasks[0],
		domain.Tasks[1],
		domain.Tasks[2],
	}
	fmt.Println("Collection instance created!")
	Collections.DeleteMany(context.TODO(), bson.D{{}})
	Collections.InsertMany(context.TODO(), tasks)
}
func ConnectDB() *mongo.Database {
	var Database *mongo.Database

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var db_url = os.Getenv("DB_URI")
	defer func() {

		fmt.Println(db_url)
	}()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(db_url).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {

		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Database = client.Database("task_manager")
	return Database
}

// Task struct
