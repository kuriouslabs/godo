package controllers

type Error struct {
	err string
}

var (
	ErrBadRequest = NewError("Bad Request")
)

func NewError(msg string) *Error {
	return &Error{
		err: msg,
	}
}

func (e Error) Error() string {
	return e.err
}
