package main

import (
	"go_sample/database"
	"go_sample/models"
	"go_sample/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// DB接続
	database.Connect()
	database.DB.AutoMigrate(&models.ToDo{})

	// ルーティング
	r.GET("/todos", routes.GetToDos)
	r.GET("/todos/:id", routes.GetToDoByID)
	r.POST("/todos", routes.CreateToDo)
	r.PUT("/todos/:id", routes.UpdateToDo)
	r.DELETE("/todos/:id", routes.DeleteToDo)

	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
