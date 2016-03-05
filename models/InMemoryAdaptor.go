package models

type InMemoryAdaptor struct {
	todoStore map[string]Todo
}

func NewInMemoryAdaptor() *InMemoryAdaptor {
	return &InMemoryAdaptor{
		todoStore: make(map[string]Todo),
	}
}

//
//// Todos
//

func (a *InMemoryAdaptor) TodoByID(id string) (*Todo, error) {
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

func (a *InMemoryAdaptor) SaveTodo(todo Todo) {
	a.todoStore[todo.ID] = todo
}

func (a *InMemoryAdaptor) DeleteTodo(todo Todo) {
	delete(a.todoStore, todo.ID)
}
