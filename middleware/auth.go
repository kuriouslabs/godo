package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/repos"
	"github.com/kuriouslabs/godo/util"
)

// Authenticated middleware that checks to see if the token is valid
func Authenticated(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := repos.TokenFromRequest(r)

		switch err {
		case nil:
			if uid, ok := repos.UserIDFromToken(token); ok {
				util.AppendUserIDToRequest(r, uid)
				h(w, r, ps)
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		case repos.ErrTokenIsInvalid,
			repos.ErrTokenIsMissing,
			repos.ErrTokenIsExpired:
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
