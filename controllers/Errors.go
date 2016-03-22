package controllers

type Error struct {
	err string
}

var (
	ErrBadRequest           = NewError("Bad Request")
	ErrUnauthorized         = NewError("Not authorized")
	ErrEntityCreationFailed = NewError("Creation of entity failed")
	ErrInternalError        = NewError("Internal error")
	ErrInvalidToken         = NewError("Invalid token")
)

func NewError(msg string) *Error {
	return &Error{
		err: msg,
	}
}

func (e Error) Error() string {
	return e.err
}
