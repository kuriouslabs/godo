package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/controllers"
	"gopkg.in/unrolled/render.v1"
)

// DataHandler is a function which takes a http request and returns an object
// and a status code. The Respond function will take these values and render them
// based on the http request accept headers.
type DataHandler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, int)

// Respond is a function which wraps a DataHandler and converts it to a httprouter.Handle function
func Respond(render *render.Render, dh controllers.Action) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result := dh(w, r, ps)
		if err := result.Error; err != nil {
			//TODO: Map error codes
			render.JSON(w, mapError(err), "")
		} else {
			render.JSON(w, 200, result.Value)
		}
	}
}

func mapError(err *controllers.Error) int {
	switch err {
	case controllers.ErrFIXME:
		return 303
	default:
		return 500
	}
}
