package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDoTitle(t *testing.T) {
  todo := ToDo{
    Title:       "Test Title",
    Description: "Test Description",
    Completed:   false,
  }
  
  assert.Equal(t, "Test Title", todo.Title, "The title should be 'Test Title'")
  assert.Equal(t, "Test Description", todo.Description, "The description should be 'Test Description'")
  assert.Equal(t, false, todo.Completed, "The completed status should be false")
}
