package models

import (
	"testing"
)

func TestById_ReturnsCorrectErrorForMissingTodo(t *testing.T) {
	repo := NewTodoRepo()
	todo, err := repo.ById("some invalid id")

	if todo != nil {
		t.Errorf("should not return a Todo: got %v", todo)
	}

	if _, ok := err.(*TodoNotFoundError); !ok {
		t.Errorf("Returned wrong error type: got %v", err)
	}
}

func TestSave_actuallySaves(t *testing.T) {
	repo := NewTodoRepo()
	todo := NewTodo("save me")

	if tmpTodo, _ := repo.ById(todo.Id); tmpTodo != nil {
		t.Error("Invalid test... todo already exists")
	}

	repo.Save(todo)

	if tmpTodo, _ := repo.ById(todo.Id); tmpTodo == nil {
		t.Errorf("Unable to save todo with id '%s'", todo.Id)
	}
}

func TestSaveAndById_returnsEqualObjects(t *testing.T) {
	repo := NewTodoRepo()
	todo := NewTodo("save me")

	repo.Save(todo)

	if tmpTodo, _ := repo.ById(todo.Id); !todo.Equal(tmpTodo) {
		t.Error("Save and ById should return Equal object")
	}
}

func TestDelete_actuallyDeletes(t *testing.T) {
	repo := NewTodoRepo()
	todo := NewTodo("save me")

	repo.Save(todo)

	if tmpTodo, _ := repo.ById(todo.Id); tmpTodo == nil {
		t.Error("Invalid test... should exist in the repo before trying to delete")
	}

	repo.Delete(todo)

	if tmpTodo, _ := repo.ById(todo.Id); tmpTodo != nil {
		t.Error("Failed to delete todo")
	}
}

func TestContains(t *testing.T) {
	repo := NewTodoRepo()
	t1 := NewTodo("t1")
	t2 := NewTodo("t2")

	repo.Save(t1)

	if repo.Contains(t2.Id) {
		t.Error("Should not contain t2")
	}

	if !repo.Contains(t1.Id) {
		t.Error("Should contain t1")
	}
}
