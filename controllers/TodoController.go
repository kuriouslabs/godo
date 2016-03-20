package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/models"
	"github.com/kuriouslabs/godo/util"
)

// TodoController represents a type of controller which is
// instantiated with a TodoRepo
type TodoController struct {
	Controller
}

// NewTodoController is a convenience function for creating a new todo controller
func NewTodoController() *TodoController {
	return &TodoController{
		Controller: NewController(),
	}
}

// Show returns a single Todo item
func (c *TodoController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	id := ps.ByName("id")
	v, _ := c.env.TodoRepo.ByID(id)
	//TODO: make this work correctly
	return Succeed(v)
}

// Create creates a new Todo item
func (c *TodoController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	uid := util.GetUserIDFromRequest(r)
	v := util.NewValidator(r)
	todoValue := v.String("todo")

	return c.AfterValidation(v, func() Result {
		t := models.NewTodo(todoValue, uid)

		c.env.TodoRepo.Save(t)

		return Succeed(t)
	})
}
