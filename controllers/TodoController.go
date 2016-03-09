package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/models"
	"net/http"
)

// TodoController represents a type of controller which is
// instantiated with a TodoRepo
type TodoController struct {
	todoRepo *models.TodoRepo
}

// NewTodoController is a convenience function for creating a
// new TodoController with a given TodoRepo
func NewTodoController(repo *models.TodoRepo) *TodoController {
	return &TodoController{
		todoRepo: repo,
	}
}

// Show returns a single Todo item
func (c *TodoController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, int) {
	id := ps.ByName("id")
	v, _ := c.todoRepo.ByID(id)
	//TODO: make this work correctly
	return v, 200
}

// Create creates a new Todo item
func (c *TodoController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, int) {
	v := r.FormValue("todo")
	t := models.NewTodo(v, "user-id-here")

	c.todoRepo.Save(t)

	return t, 200
}
