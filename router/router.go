package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
	"github.com/zaahidali/task_manager_api/middleware"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", middleware.AuthMiddleware(), controllers.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware(), controllers.GetTasksId)
	router.POST("/tasks", middleware.AuthMiddleware(), controllers.CreateTask)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), controllers.UpdateTask)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), controllers.DeleteTask)

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
