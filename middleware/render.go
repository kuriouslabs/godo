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
		if result.Error != nil {
			statusCode := mapError(result.Error)
			errMessage := map[string]string{
				"error":  http.StatusText(statusCode),
				"reason": result.Reason,
			}

			render.JSON(w, statusCode, errMessage)
		} else {
			render.JSON(w, http.StatusOK, result.Value)
		}
	}
}

func mapError(err *controllers.Error) int {
	switch err {
	case controllers.ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
