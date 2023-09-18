package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"rest_api/internal/app/handler"
	"rest_api/internal/app/repository/task"
	user "rest_api/internal/app/repository/user"

	_ "github.com/go-sql-driver/mysql"

	"rest_api/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DB")
	// Create the MySQL connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB)
	// Connect to the MySQL database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()
	r := gin.Default()
	fmt.Printf("Connected to the database! %v ", db)
	//* TASK
	taskRepo := task.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	// Initialize handlers
	taskHandler := handler.NewTaskHandler(*taskService)
	//* USER
	userRepo := user.NewMySQLUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(*userService)

	// Define routes
	r.GET("/tasks", taskHandler.ListTasks)
	r.POST("/tasks", taskHandler.AddTask)
	r.POST("/users/create", userHandler.CreateUser)

	// Start the server
	r.Run(":8080")
}
