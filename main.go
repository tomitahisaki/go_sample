package main

import (
	"go_sample/api/routes"
	database "go_sample/config"
	"go_sample/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// DB接続
	database.Connect()
	database.DB.AutoMigrate(&models.ToDo{})

	// ルーター設定
	routes.RegisterToDoRoutes(r)

	// サーバー起動
	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
