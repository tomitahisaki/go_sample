package handlers

import (
	"go_sample/models"
	"go_sample/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetToDos(c *gin.Context) {
	todos, err := services.GetToDos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetToDoByID(c *gin.Context) {
  id := c.Param("id")
  todo, err := services.GetToDoByID(id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
    return
  }
  c.JSON(http.StatusOK, todo)
}


func CreateToDo(c *gin.Context) {
	var todo models.ToDo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := services.AddToDo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusCreated, todo)
}
