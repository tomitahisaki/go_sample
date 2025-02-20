package services

import (
	"go_sample/models"
	"go_sample/repository"
)

func GetToDos() ([]models.ToDo, error) {
	return repository.GetAllToDos()
}

func GetToDoByID(id string) (models.ToDo, error) {
  return repository.GetToDoByID(id)
}

func AddToDo(todo *models.ToDo) error {
	return repository.CreateToDo(todo)
}
