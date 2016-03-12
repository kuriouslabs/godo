package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/models"
	"github.com/kuriouslabs/godo/repos"
	"github.com/kuriouslabs/godo/util"
)

// TodoController represents a type of controller which is
// instantiated with a TodoRepo
type TodoController struct {
	todoRepo *repos.TodoRepo
}

// NewTodoController is a convenience function for creating a
// new TodoController with a given TodoRepo
func NewTodoController(repo *repos.TodoRepo) *TodoController {
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
	uid := util.GetUserIDFromRequest(r)
	v := r.FormValue("todo")
	t := models.NewTodo(v, uid)

	c.todoRepo.Save(t)

	return t, 200
}
