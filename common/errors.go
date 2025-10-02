package common

import (
	"errors"
	"time"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

var (
	ErrNotFound = errors.New("record not found")
	ErrConflict = errors.New("resource already exists")

	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrDuplicatePhone    = errors.New("phone already exists")

	QueryTimeoutDuration = time.Second * 5
)
