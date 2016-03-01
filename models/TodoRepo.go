package models

import (
	"fmt"
)

/// Errors

type TodoNotFoundError struct {
	Id string
}

func (e *TodoNotFoundError) Error() string {
	return fmt.Sprintf("Cannot find Todo with id: '%s'", e.Id)
}

/// TodoRepo

type TodoRepo struct {
	_store map[string]Todo
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{
		_store: make(map[string]Todo),
	}
}

/// Query Methods

func (t *TodoRepo) ById(id string) (*Todo, error) {

	todo, ok := t._store[id]
	if ok {
		return &todo, nil
	}
	return nil, &TodoNotFoundError{id}
}

func (t *TodoRepo) Contains(id string) bool {
	_, ok := t._store[id]
	return ok
}

/// Saving

func (t *TodoRepo) Save(todo Todo) {
	t._store[todo.Id] = todo
}

/// Deleting

func (t *TodoRepo) Delete(todo Todo) {
	delete(t._store, todo.Id)
}
