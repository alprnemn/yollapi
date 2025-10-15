package common

import (
	"errors"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("resource already exists")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrDuplicatePhone    = errors.New("phone already exists")

	ErrGeneratePassword = errors.New("error occurred while generating password")
	ErrPasswordInvalid  = errors.New("invalid password")

	ErrMissingAuthHeader = errors.New("missing Authorization header")
	ErrWrongAuthHeader   = errors.New("wrong Authorization header value")
)
