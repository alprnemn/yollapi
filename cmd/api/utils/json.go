package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alprnemn/yollapi/common"
	"io"
	"net/http"
)

const maxBytes = 1 << 20

func ParseJSON(w http.ResponseWriter, req *http.Request, data any) error {

	req.Body = http.MaxBytesReader(w, req.Body, int64(maxBytes))

	if req.Body == nil {
		return fmt.Errorf("missing request body")
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	// Decode into the provided data
	if err := decoder.Decode(data); err != nil {
		if errors.Is(err, io.EOF) {
			return fmt.Errorf("request body must not be empty")
		}

		// Wrong type (e.g., string instead of number)
		var ute *json.UnmarshalTypeError
		if errors.As(err, &ute) {
			return fmt.Errorf("field '%s' must be of type %s", ute.Field, ute.Type.String())
		}

		// Syntax error
		var se *json.SyntaxError
		if errors.As(err, &se) {
			return fmt.Errorf("badly-formed JSON at position %d", se.Offset)
		}

		return err
	}

	// Ensure only one JSON object
	if decoder.More() {
		return fmt.Errorf("request body must only contain a single JSON object")
	}

	return nil
}

func WriteError(w http.ResponseWriter, status int, message string) {
	resp := common.ErrorResponse{
		Code:    status,
		Message: message,
	}
	_ = WriteJSON(w, status, resp)
}

func JsonResponse(w http.ResponseWriter, status int, data any) error {
	return WriteJSON(w, status, data)
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
