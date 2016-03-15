package controllers

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/models"
	"github.com/kuriouslabs/godo/repos"
)

// AuthController a controller for handling auth
type AuthController struct {
	tokenRepo *repos.TokenRepo
}

// NewAuthController creates a new auth controller
func NewAuthController() *AuthController {
	return &AuthController{
		tokenRepo: nil,
	}
}

// LogIn attempt to log in
func (c *AuthController) LogIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, int) {
	// for now just log the user in that was passed in
	uid := r.FormValue("user_id")

	if uid == "" {
		// TODO: which error code should we return
		return nil, http.StatusForbidden
	}

	// TODO: validate login

	exp := time.Now().Add(time.Second * time.Duration(60))
	token := repos.GenerateTokenForUser(uid, exp)

	//TODO: use the UserRepo
	user := models.NewUser("user123", uid)

	return map[string]interface{}{
		"token": token,
		"user":  user,
	}, http.StatusOK
}
