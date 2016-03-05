package models

import (
	"github.com/satori/go.uuid"
)

// Todo object
type Todo struct {
	ID        string `json:"id"`
	Value     string `json:"name"`
	Completed bool   `json:"completed"`
	UserID    string `json:"user_id"`
}

// NewTodo returns a new Todo object
func NewTodo(value string, userID string) Todo {
	return Todo{
		ID:        uuid.NewV4().String(),
		Value:     value,
		Completed: false,
		UserID:    userID,
	}
}

// Equal returns true if the todo objects are the same
func (t *Todo) Equal(other *Todo) bool {
	if other == nil {
		return false
	}

	return t.ID == other.ID &&
		t.Value == other.Value &&
		t.Completed == other.Completed &&
		t.UserID == other.UserID
}
