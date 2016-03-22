package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kuriouslabs/godo/repos"
	"github.com/kuriouslabs/godo/util"
)

// Authenticated wrapper that checks to see if the token is valid
func Authenticated(h Action) Action {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) Result {
		token, err := repos.TokenFromRequest(r)

		switch err {
		case nil:
			if uid, ok := repos.UserIDFromToken(token); ok {
				util.AppendUserIDToRequest(r, uid)
				return h(w, r, ps)
			}
			return Fail(ErrInternalError, "Error extracting user id from token")
		case repos.ErrTokenIsInvalid:
			return Fail(ErrInvalidToken, "Provided token is invalid")
		case repos.ErrTokenIsMissing:
			return Fail(ErrInvalidToken, "Required token is missing")
		case repos.ErrTokenIsExpired:
			return Fail(ErrInvalidToken, "Provided token is expired")
		}

		return Fail(ErrInternalError, "Unknown server error")
	}
}
