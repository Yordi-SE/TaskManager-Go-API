package controllers

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	GetTasks() func(c *gin.Context)
	GetTasksId() func(c *gin.Context)
	CreateTask() func(c *gin.Context)
	UpdateTask() func(c *gin.Context)
	DeleteTask() func(c *gin.Context)
}

type AuthControllerInterface interface {
	Register() func(c *gin.Context)
	Login() func(c *gin.Context)
	Promote() func(c *gin.Context)
}
