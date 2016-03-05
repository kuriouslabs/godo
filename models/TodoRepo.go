package models

import (
	"fmt"
)

// TodoRepoAdaptor ...
// Provides a interface for adaptors for the TodoRepo
type TodoRepoAdaptor interface {
	TodoByID(id string) (*Todo, error)
	ContainsTodo(id string) bool
	SaveTodo(todo Todo)
	DeleteTodo(todo Todo)
}

/// Errors

type TodoNotFoundError struct {
	Id string
}

func (e *TodoNotFoundError) Error() string {
	return fmt.Sprintf("Cannot find Todo with id: '%s'", e.Id)
}

/// TodoRepo
type TodoRepo struct {
	adaptor TodoRepoAdaptor
}

// NewTodoRepo instantiates a new repo
func NewTodoRepo(adaptor TodoRepoAdaptor) *TodoRepo {
	return &TodoRepo{
		adaptor: adaptor,
	}
}

// ByID returns a Todo with the given ID
func (t *TodoRepo) ByID(id string) (*Todo, error) {
	return t.adaptor.TodoByID(id)
}

// Contains returns true if the todo exists in the repo
func (t *TodoRepo) Contains(id string) bool {
	return t.adaptor.ContainsTodo(id)
}

// Save saves the current todo
func (t *TodoRepo) Save(todo Todo) {
	t.adaptor.SaveTodo(todo)
}

// Delete deletes the current Todo
func (t *TodoRepo) Delete(todo Todo) {
	t.adaptor.DeleteTodo(todo)
}
