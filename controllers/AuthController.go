package controllers

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/repos"
	"github.com/kuriouslabs/godo/util"
)

// AuthController a controller for handling auth
type AuthController struct {
	Controller
}

// NewAuthController creates a new auth controller
func NewAuthController() *AuthController {
	return &AuthController{
		Controller: NewController(),
	}
}

// LogIn attempt to log in
func (c *AuthController) LogIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	// for now just log the user in that was passed in
	v := util.NewValidator(r)
	uid := v.String("user_id")
	pw := v.String("password")

	return c.AfterValidation(v, func() Result {
		if !c.env.UserRepo.AuthenticateUserPassword(uid, pw) {
			return Fail(ErrUnauthorized, "invalid username or password")
		}

		exp := time.Now().Add(time.Hour * time.Duration(72))
		token := repos.GenerateTokenForUser(uid, exp)

		user, _ := c.env.UserRepo.ByID(uid)

		return Succeed(map[string]interface{}{
			"token": token,
			"user":  user,
		})
	})
}
