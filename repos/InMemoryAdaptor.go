package repos

import "github.com/kuriouslabs/godo/models"

type InMemoryAdaptor struct {
	todoStore map[string]models.Todo
}

func NewInMemoryAdaptor() *InMemoryAdaptor {
	return &InMemoryAdaptor{
		todoStore: make(map[string]models.Todo),
	}
}

//
//// Todos
//

func (a *InMemoryAdaptor) TodoByID(id string) (*models.Todo, error) {
	todo, ok := a.todoStore[id]
	if ok {
		return &todo, nil
	}

	return nil, &TodoNotFoundError{id}
}

func (a *InMemoryAdaptor) ContainsTodo(id string) bool {
	_, ok := a.todoStore[id]
	return ok
}

func (a *InMemoryAdaptor) SaveTodo(todo models.Todo) {
	a.todoStore[todo.ID] = todo
}

func (a *InMemoryAdaptor) DeleteTodo(todo models.Todo) {
	delete(a.todoStore, todo.ID)
}
