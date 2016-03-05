package models

import (
	"testing"
)

func TestNewTodoValuesSet(t *testing.T) {
	v := "Do something"
	uid := "123"
	todo := NewTodo(v, uid)

	if todo.Value != v || todo.Completed == true || todo.UserID != uid {
		t.Error("Failed to set proper values on NewTodo")
	}
}

func TestNewTodoUniqueId(t *testing.T) {
	t1 := NewTodo("first", "123")
	t2 := NewTodo("second", "123")
	if t1.ID == t2.ID {
		t.Errorf("NewTodo should generate unique Id got %s, %s", t1.ID, t2.ID)
	}
}

func TestTodoEquality(t *testing.T) {
	id := "123"
	v := "value"
	c := false

	t1 := &Todo{ID: id, Value: v, Completed: c}
	t2 := &Todo{ID: id, Value: v, Completed: c}

	if !t1.Equal(t2) {
		t.Error("Values should be equal")
	}
}

func TestTodoEquality_nilOther(t *testing.T) {
	id := "123"
	v := "value"
	c := false

	t1 := &Todo{ID: id, Value: v, Completed: c}
	var t2 *Todo

	if t1.Equal(t2) {
		t.Error("Should not be equal to nil value")
	}
}

func TestTodoEquality_failsForCompletion(t *testing.T) {
	id := "123"
	v := "value"
	uid := "userID"

	t1 := &Todo{ID: id, Value: v, Completed: true, UserID: uid}
	t2 := &Todo{ID: id, Value: v, Completed: false, UserID: uid}

	if t1.Equal(t2) {
		t.Error("Values should not be equal if Completed is different")
	}
}

func TestTodoEquality_failsForUserID(t *testing.T) {
	id := "123"
	v := "value"

	t1 := &Todo{ID: id, Value: v, Completed: true, UserID: "user_id_1"}
	t2 := &Todo{ID: id, Value: v, Completed: true, UserID: "user_id_2"}

	if t1.Equal(t2) {
		t.Error("Values should not be equal if UserID is different")
	}
}
