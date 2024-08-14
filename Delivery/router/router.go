package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/Delivery/controllers"
	"github.com/zaahidali/task_manager_api/Infrastructure"
)

type RouterHandler struct {
	TaskController *controllers.TaskHandler
	AuthController *controllers.AuthHandler
}

var Route = RouterHandler{
	TaskController: new(controllers.TaskHandler),
	AuthController: new(controllers.AuthHandler),
}

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", Infrastructure.AuthMiddleware(), Route.TaskController.GetTasks())
	router.GET("/tasks/:id", Infrastructure.AuthMiddleware(), Route.TaskController.GetTasksId())
	router.POST("/tasks", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), Route.TaskController.CreateTask())
	router.PUT("/tasks/:id", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), Route.TaskController.UpdateTask())
	router.DELETE("/tasks/:id", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), Route.TaskController.DeleteTask())
	router.POST("/promote", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), Route.AuthController.Promote())

	// auth routes
	router.POST("/auth", Route.AuthController.Register())
	router.POST("/auth/login", Route.AuthController.Login())

	return router
}

// Run starts the router
func Run() {
	router := SetupRouter()
	router.Run()
}
