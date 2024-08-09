package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/Delivery/controllers"
	"github.com/zaahidali/task_manager_api/Infrastructure"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", Infrastructure.AuthMiddleware(), controllers.GetTasks)
	router.GET("/tasks/:id", Infrastructure.AuthMiddleware(), controllers.GetTasksId)
	router.POST("/tasks", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), controllers.CreateTask)
	router.PUT("/tasks/:id", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), controllers.UpdateTask)
	router.DELETE("/tasks/:id", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), controllers.DeleteTask)
	router.POST("/promote", Infrastructure.AuthMiddleware(), Infrastructure.Isadmin(), controllers.Promote)

	// auth routes
	router.POST("/auth", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	return router
}

// Run starts the router
func Run() {
	router := SetupRouter()
	router.Run()
}
