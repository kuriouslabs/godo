package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Result struct {
	Value  interface{}
	Error  *Error
	Reason string
}

type Action func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result

func Succeed(val interface{}) Result {
	return Result{
		Value: val,
	}
}

func Fail(err *Error, reason string) Result {
	return Result{
		Error:  err,
		Reason: reason,
	}
}
