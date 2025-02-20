package repository

import (
	database "go_sample/config"
	"go_sample/models"
)

// GetAllToDos is a function to get all todos
func GetAllToDos() ([]models.ToDo, error) {
	var todos []models.ToDo
	result := database.DB.Find(&todos)
	return todos, result.Error
}

// GetToDoByID is a function to get a todo by ID
func GetToDoByID(id string) (models.ToDo, error) {
	var todo models.ToDo
	result := database.DB.First(&todo, id)
	return todo, result.Error
}

// createToDo is a function to create a new todo
func CreateToDo(todo *models.ToDo) error {
	result := database.DB.Create(todo)
	return result.Error
}

