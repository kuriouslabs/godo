package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Authenticated(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//TODO: Implement the authentication stack
		authenticated := true
		if authenticated {
			h(w, r, ps)
			return
		}

		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
