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

func TestTodoEquality(t *testing.T) {
	id := "123"
	v := "value"
	c := false

	t1 := &Todo{Id: id, Value: v, Completed: c}
	t2 := &Todo{Id: id, Value: v, Completed: c}

	if !t1.Equal(t2) {
		t.Error("Values should be equal")
	}
}

func TestTodoEquality_nilOther(t *testing.T) {
	id := "123"
	v := "value"
	c := false

	t1 := &Todo{Id: id, Value: v, Completed: c}
	var t2 *Todo = nil

	if t1.Equal(t2) {
		t.Error("Should not be equal to nil value")
	}
}

func TestTodoEquality_failsForCompletion(t *testing.T) {
	id := "123"
	v := "value"

	t1 := &Todo{Id: id, Value: v, Completed: true}
	t2 := &Todo{Id: id, Value: v, Completed: false}

	if t1.Equal(t2) {
		t.Error("Values should not be equal if Completed is different")
	}
}
