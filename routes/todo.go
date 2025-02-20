package routes

import (
	"go_sample/database"
	"go_sample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ToDo一覧を取得
func GetToDos(c *gin.Context) {
	var todos []models.ToDo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// ToDoをIDで取得
func GetToDoByID(c *gin.Context) {
	id := c.Param("id")
	var todo models.ToDo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// ToDoを作成
func CreateToDo(c *gin.Context) {
	var todo models.ToDo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// ToDoを更新
func UpdateToDo(c *gin.Context) {
	id := c.Param("id")
	var todo models.ToDo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// ToDoを削除
func DeleteToDo(c *gin.Context) {
	id := c.Param("id")
	var todo models.ToDo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "ToDo deleted"})
}
