package repos

import (
	"errors"

	"github.com/kuriouslabs/godo/models"
)

// TodoRepoAdaptor ...
// Provides a interface for adaptors for the TodoRepo
type TodoRepoAdaptor interface {
	TodoByID(id string) (*models.Todo, error)
	ContainsTodo(id string) bool
	SaveTodo(todo models.Todo)
	DeleteTodo(todo models.Todo)
}

var (
	ErrTodoNotFound = errors.New("Cannot find todo")
)

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
func (t *TodoRepo) ByID(id string) (*models.Todo, error) {
	return t.adaptor.TodoByID(id)
}

// Contains returns true if the todo exists in the repo
func (t *TodoRepo) Contains(id string) bool {
	return t.adaptor.ContainsTodo(id)
}

// Save saves the current todo
func (t *TodoRepo) Save(todo models.Todo) {
	t.adaptor.SaveTodo(todo)
}

// Delete deletes the current Todo
func (t *TodoRepo) Delete(todo models.Todo) {
	t.adaptor.DeleteTodo(todo)
}
