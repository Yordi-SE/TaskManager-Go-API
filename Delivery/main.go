package main

import (
	"fmt"

	"github.com/zaahidali/task_manager_api/Delivery/router"
	repositories "github.com/zaahidali/task_manager_api/Repositories"
)

func main() {
	router.Run()
}

func init() {
	repositories.Database = repositories.ConnectDB()
	repositories.Collections = repositories.Database.Collection("tasks")
	repositories.UserCollection = repositories.Database.Collection("User")
	fmt.Println("Collection instance created!")
}
