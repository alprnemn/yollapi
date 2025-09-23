package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
		return err
	}

	// Ensure only one JSON object
	if decoder.More() {
		return fmt.Errorf("request body must only contain a single JSON object")
	}

	return nil
}

func WriteError(w http.ResponseWriter, status int, message string) {
	type envelope struct {
		Error string `json:"error"`
	}
	_ = WriteJSON(w, status, envelope{Error: message})
}

func JsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}
	return WriteJSON(w, status, envelope{Data: data})
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
