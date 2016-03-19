package controllers

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/config"
	"github.com/kuriouslabs/godo/repos"
)

// AuthController a controller for handling auth
type AuthController struct {
	env *config.Env
}

// NewAuthController creates a new auth controller
func NewAuthController(env *config.Env) *AuthController {
	return &AuthController{
		env: env,
	}
}

// LogIn attempt to log in
func (c *AuthController) LogIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	// for now just log the user in that was passed in
	uid := r.FormValue("user_id")
	pw := r.FormValue("password")

	//TODO: figure out a better validation scheme
	if uid == "" {
		return Fail(ErrBadRequest, "missing parameter user_id")
	}

	if pw == "" {
		return Fail(ErrBadRequest, "missing parameter password")
	}

	if !c.env.UserRepo.AuthenticateUserPassword(uid, pw) {
		return Fail(ErrUnauthorized, "invalid username or password")
	}

	exp := time.Now().Add(time.Second * time.Duration(60))
	token := repos.GenerateTokenForUser(uid, exp)

	user, _ := c.env.UserRepo.ByID(uid)

	return Succeed(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}
