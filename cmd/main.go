package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"rest_api/internal/app/handler"
	repository "rest_api/internal/app/repository/task"

	"rest_api/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createMongoClient() (*mongo.Client, error) {
	// Define the MongoDB connection string
	// connectionString := "mongodb://localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://Collins01:1223334444@nodeexpressprojects.xkvg6tw.mongodb.net/03-TASK-MANAGER-APP?retryWrites=true&w=majority")
	// Create a MongoDB client
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func accessDatabase(client *mongo.Client) {
	// Get a handle to the database
	// database := client.Database("mydatabase")

	// Now you can work with the "mydatabase" database
	// For example, you can access collections, insert data, query data, etc.
}

func closeMongoClient(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		fmt.Println("Error disconnecting from MongoDB:", err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	r := gin.Default()

	// // Initialize MongoDB client
	// client, err := createMongoClient()
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Disconnect(context.Background())
	// // Initialize repository and service
	// Create a MySQL database connection
	db, err := sql.Open("mysql", "root:mysecretpassword@tcp(localhost:3306)/mydb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Check the database connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to the database!")
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
