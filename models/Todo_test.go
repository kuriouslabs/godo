package models

import (
	"testing"
)

func TestNewTodoValuesSet(t *testing.T) {
	v := "Do something"
	todo := NewTodo(v)

	if todo.Value != v || todo.Completed == true {
		t.Error("Failed to set proper values on NewTodo")
	}
}

func TestNewTodoUniqueId(t *testing.T) {
	t1 := NewTodo("first")
	t2 := NewTodo("second")

	if t1.Id == t2.Id {
		t.Errorf("NewTodo should generate unique Id got %s, %s", t1.Id, t2.Id)
	}
}
