package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/models"
	"github.com/kuriouslabs/godo/util"
)

// UserController a controller for mananging users
type UserController struct {
	Controller
}

// NewUserController creates a new user controller
func NewUserController() *UserController {
	return &UserController{
		Controller: NewController(),
	}
}

// Create creates the given user
func (c *UserController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	v := util.NewValidator(r)
	name := v.String("name")
	username := v.String("username")

	return c.AfterValidation(v, func() Result {
		user := models.NewUser(name, username)
		if err := c.env.UserRepo.CreateUser(&user); err != nil {
			//TODO: better error
			return Fail(ErrEntityCreationFailed, "Create user failed")
		}
		return Succeed(user)
	})
}

// Me returns the current logged in user
func (c *UserController) Me(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	uid := util.GetUserIDFromRequest(r)
	user, err := c.env.UserRepo.ByID(uid)
	if err != nil {
		//TODO: Wrong error
		return Fail(ErrBadRequest, "User not found")
	}
	return Succeed(user)
}
