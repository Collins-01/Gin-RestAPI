package main

import (
	"rest_api/internal/app/handler"
	repository "rest_api/internal/app/repository/task"
	"rest_api/internal/app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize repository and service
	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	// Initialize handlers
	taskHandler := handler.NewTaskHandler(*taskService)

	// Define routes
	r.GET("/tasks", taskHandler.ListTasks)
	r.POST("/tasks", taskHandler.AddTask)

	// Start the server
	r.Run(":8080")
}
