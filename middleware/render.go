package middleware

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

// DataHandler is a function which takes a http request and returns an object
// and a status code. The Respond function will take these values and render them
// based on the http request accept headers.
type DataHandler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, int)

// Respond is a function which wraps a DataHandler and converts it to a httprouter.Handle function
func Respond(render *render.Render, dh DataHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		value, statusCode := dh(w, r, ps)
		if value != nil {
			render.JSON(w, statusCode, value)
		} else {
			render.JSON(w, statusCode, "")
		}
	}
}
