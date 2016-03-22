package controllers

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/models"
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
	username := v.String("username")
	pw := v.String("password")

	return c.AfterValidation(v, func() Result {
		user := c.env.UserRepo.AuthenticateUserPassword(username, pw)
		if user == nil {
			return Fail(ErrUnauthorized, "invalid username or password")
		}

		token, exp := generateToken(user)
		refreshToken := c.env.TokenRepo.GenerateRefreshTokenForUser(user.ID)

		return Succeed(map[string]interface{}{
			"token":   token,
			"user":    user,
			"exp":     exp,
			"refresh": refreshToken,
		})
	})
}

func (c *AuthController) LogOut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
	uid := util.GetUserIDFromRequest(r)

	c.env.TokenRepo.RevokeRefreshTokenForUser(uid)
	//TODO: Revoke the current jwt-token as well?
	return Succeed("")
}

// RefreshToken refreshes the given token
func (c *AuthController) RefreshToken(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) Result {
	v := util.NewValidator(r)
	refreshToken := v.String("refresh_token")
	uid := v.String("user_id")

	return c.AfterValidation(v, func() Result {
		if !c.env.TokenRepo.ValidateRefreshToken(refreshToken, uid) {
			return Fail(ErrUnauthorized, "invalid refresh token")
		}

		user, err := c.env.UserRepo.ByID(uid)

		if err != nil {
			return Fail(ErrBadRequest, "Could not find user for refresh token")
		}

		token, exp := generateToken(user)

		return Succeed(map[string]interface{}{
			"token": token,
			"user":  user,
			"exp":   exp,
		})
	})
}

func generateToken(user *models.User) (string, time.Time) {
	exp := time.Now().Add(time.Hour * time.Duration(1))
	token := repos.GenerateTokenForUser(user.ID, exp)
	return token, exp
}
