package domain

import (
	"errors"
)

// will contain common errors
var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrBadRequest         = errors.New("bad request")
	ErrValidation         = errors.New("validation error")
	ErrSomethingWentWrong = errors.New("something went wrong, please try again")
	ErrCacheSet           = errors.New("failed to set cache")
)

// CustomError defines custom error
type CustomError struct {
	Message        string
	HttpStatusCode int
}

// Error returns error message
func (e *CustomError) Error() string {
	return e.Message
}

// StatusCode returns http status code of err
func (e *CustomError) StatusCode() int {
	return e.HttpStatusCode
}
