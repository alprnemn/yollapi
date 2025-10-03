package userservice

import (
	"errors"
	"net/http"

	cmn "github.com/alprnemn/yollapi/common"
)

func handleErrorRegister(err error) *cmn.ErrorResponse {
	// handle repository error as conflict or internal server error
	switch {
	case errors.Is(err, cmn.ErrDuplicateUsername),
		errors.Is(err, cmn.ErrDuplicateEmail),
		errors.Is(err, cmn.ErrDuplicatePhone):
		return &cmn.ErrorResponse{Code: http.StatusConflict, Message: err.Error()}
	}
	return &cmn.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()}
}
