package models

import (
	"testing"
)

func makeRepo() *TodoRepo {
	adaptor := NewInMemoryAdaptor()
	repo := NewTodoRepo(adaptor)
	return repo
}

func TestById_ReturnsCorrectErrorForMissingTodo(t *testing.T) {
	repo := makeRepo()
	todo, err := repo.ByID("some invalid id")

	if todo != nil {
		t.Errorf("should not return a Todo: got %v", todo)
	}

	if _, ok := err.(*TodoNotFoundError); !ok {
		t.Errorf("Returned wrong error type: got %v", err)
	}
}

func TestSave_actuallySaves(t *testing.T) {
	repo := makeRepo()
	todo := NewTodo("save me", "123")

	if tmpTodo, _ := repo.ByID(todo.ID); tmpTodo != nil {
		t.Error("Invalid test... todo already exists")
	}

	repo.Save(todo)

	if tmpTodo, _ := repo.ByID(todo.ID); tmpTodo == nil {
		t.Errorf("Unable to save todo with id '%s'", todo.ID)
	}
}

func TestSaveAndById_returnsEqualObjects(t *testing.T) {
	repo := makeRepo()
	todo := NewTodo("save me", "123")

	repo.Save(todo)

	if tmpTodo, _ := repo.ByID(todo.ID); !todo.Equal(tmpTodo) {
		t.Error("Save and ById should return Equal object")
	}
}

func TestDelete_actuallyDeletes(t *testing.T) {
	repo := makeRepo()
	todo := NewTodo("save me", "123")

	repo.Save(todo)

	if tmpTodo, _ := repo.ByID(todo.ID); tmpTodo == nil {
		t.Error("Invalid test... should exist in the repo before trying to delete")
	}

	repo.Delete(todo)

	if tmpTodo, _ := repo.ByID(todo.ID); tmpTodo != nil {
		t.Error("Failed to delete todo")
	}
}

func TestContains(t *testing.T) {
	repo := makeRepo()
	t1 := NewTodo("t1", "123")
	t2 := NewTodo("t2", "123")

	repo.Save(t1)

	if repo.Contains(t2.ID) {
		t.Error("Should not contain t2")
	}

	if !repo.Contains(t1.ID) {
		t.Error("Should contain t1")
	}
}
