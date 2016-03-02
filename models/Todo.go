package models

import (
	"github.com/satori/go.uuid"
)

type Todo struct {
	Id        string `json:"id"`
	Value     string `json:"name"`
	Completed bool   `json:"completed"`
}

func NewTodo(value string) Todo {
	return Todo{
		Id:        uuid.NewV4().String(),
		Value:     value,
		Completed: false,
	}
}

func (t *Todo) Equal(other *Todo) bool {
	if other == nil {
		return false
	}

	return t.Id == other.Id &&
		t.Value == other.Value &&
		t.Completed == other.Completed
}
