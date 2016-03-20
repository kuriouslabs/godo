package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/util"
)

type UserController struct {
	Controller
}

func NewUserController() *UserController {
	return &UserController{
		Controller: NewController(),
	}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	v := util.NewValidator(r)
	name := v.String("name")
	age := v.Int("age")

	return c.AfterValidation(v, func() Result {
		return Succeed(map[string]interface{}{
			"name": name,
			"age":  age,
		})
	})
}

func (c *UserController) Me(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	return Succeed("ME")
}
