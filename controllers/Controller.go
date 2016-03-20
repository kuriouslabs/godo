package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/config"
	"github.com/kuriouslabs/godo/util"
)

// Controller a base controller
type Controller struct {
	env *config.Env
}

// NewController returns a new instance of Controller
func NewController() Controller {
	if sharedEnvironment == nil {
		panic("Must have a share env")
	}
	return Controller{
		env: sharedEnvironment,
	}
}

var sharedEnvironment *config.Env

// RegisterEnv registers a shared environment
// variable. This method should be called once
// from the main thread.
func RegisterEnv(env *config.Env) {
	sharedEnvironment = env
}

func (c *Controller) AfterValidation(v *util.Validator, f func() Result) Result {
	if ok, err := v.Passed(); !ok {
		return Fail(ErrBadRequest, err.Error())
	}
	return f()
}

// Result a value to be returned by Action
type Result struct {
	Value  interface{}
	Error  *Error
	Reason string
}

// Action a function to interface with the render middleware
type Action func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result

// Succeed a convenience method to be used inside of Actions
func Succeed(val interface{}) Result {
	return Result{
		Value: val,
	}
}

// Fail a convenience method to be used inside of Actions
func Fail(err *Error, reason string) Result {
	return Result{
		Error:  err,
		Reason: reason,
	}
}
