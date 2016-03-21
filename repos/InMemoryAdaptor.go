package repos

import "github.com/kuriouslabs/godo/models"

// InMemoryAdaptor an in memory adaptor
type InMemoryAdaptor struct {
	todoStore map[string]*models.Todo
	userStore map[string]*models.User
}

// NewInMemoryAdaptor returns a new adaptor
func NewInMemoryAdaptor() *InMemoryAdaptor {
	return &InMemoryAdaptor{
		todoStore: make(map[string]*models.Todo),
		userStore: make(map[string]*models.User),
	}
}

//
//// Todos
//

// TodoByID returns a todo with the given id
func (a *InMemoryAdaptor) TodoByID(id string) (*models.Todo, error) {
	todo, ok := a.todoStore[id]
	if ok {
		return todo, nil
	}

	return nil, ErrTodoNotFound
}

// ContainsTodo checks if todo exists
func (a *InMemoryAdaptor) ContainsTodo(id string) bool {
	_, ok := a.todoStore[id]
	return ok
}

// SaveTodo saves the todo
func (a *InMemoryAdaptor) SaveTodo(todo models.Todo) {
	a.todoStore[todo.ID] = &todo
}

// DeleteTodo removes the todo
func (a *InMemoryAdaptor) DeleteTodo(todo models.Todo) {
	delete(a.todoStore, todo.ID)
}

//
//// Users
//

// UserByID returns a user with the given id
func (a *InMemoryAdaptor) UserByID(id string) (*models.User, error) {
	user, ok := a.userStore[id]
	if ok {
		return user, nil
	}

	return nil, ErrUserNotFound
}

func (a *InMemoryAdaptor) UserByUsername(username string) (*models.User, error) {
	for _, user := range a.userStore {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, ErrUserNotFound
}

// UserExists checks if user exists
func (a *InMemoryAdaptor) UserExists(id string) bool {
	_, ok := a.userStore[id]
	return ok
}

// CreateUser creates a user
func (a *InMemoryAdaptor) CreateUser(user *models.User) error {
	if a.UserExists(user.ID) {
		return ErrUserAlreadyExists
	}

	a.userStore[user.ID] = user
	return nil
}

// UpdateUser updates the user
func (a *InMemoryAdaptor) UpdateUser(user *models.User) {
	a.userStore[user.ID] = user
}
