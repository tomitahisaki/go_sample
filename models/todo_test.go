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
}
