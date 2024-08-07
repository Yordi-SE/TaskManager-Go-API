package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTasksId)
	router.POST("/tasks", controllers.CreateTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	return router
}

// Run starts the router
func Run() {
	router := SetupRouter()
	router.Run()
}
