package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Result struct {
	Value interface{}
	Error *Error
}

type Action func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result

type Error struct {
	err string
}

var (
	ErrFIXME = &Error{"FIX ME"}
)

func (e Error) Error() string {
	return e.err
}

func Succeed(val interface{}) Result {
	return Result{
		Value: val,
	}
}

func Fail(err *Error) Result {
	return Result{
		Error: err,
	}
}
