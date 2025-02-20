package routes

import (
	"go_sample/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterToDoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("", handlers.GetToDos)
    todoGroup.GET("/:id", handlers.GetToDoByID)
		todoGroup.POST("", handlers.CreateToDo)
	}
}
