package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/models"
	"net/http"
)

type TodoController struct {
	todoRepo *models.TodoRepo
}

func Todo(repo *models.TodoRepo) *TodoController {
	return &TodoController{
		todoRepo: repo,
	}
}

func (c *TodoController) Show() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println(c.todoRepo)
		fmt.Fprintf(w, "hello %s!\n", ps.ByName("name"))
	}
}
