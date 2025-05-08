package routes

import (
	"github.com/cancellap/TrabalhoFaculGolang/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	tasks := router.Group("/tasks")
	{
		tasks.POST("", handlers.CreateTask)
		tasks.GET("", handlers.GetTasks)
		tasks.PUT("/attStatus/:id", handlers.UpdateTaskStatus)
	}
}
