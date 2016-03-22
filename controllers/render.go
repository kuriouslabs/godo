package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/unrolled/render.v1"
)

// Render is a function which wraps a DataHandler and converts it to a httprouter.Handle function
func Render(render *render.Render, dh Action) httprouter.Handle {
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

func mapError(err *Error) int {
	switch err {
	case ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
