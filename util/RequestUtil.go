package util

import "net/http"

const userHeader = "user-id"

// AppendUserIDToRequest adds the specifed user id to
// the header in the given request
func AppendUserIDToRequest(r *http.Request, uid string) {
	r.Header.Add(userHeader, uid)
}

// GetUserIDFromRequest returns a string from the given
// request or nil if it doesn't exist.
func GetUserIDFromRequest(r *http.Request) string {
	return r.Header.Get(userHeader)
}
