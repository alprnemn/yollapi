package main

import "net/http"

func (app *api) healthCheckHandler(w http.ResponseWriter, req *http.Request) {

	err := WriteJSON(w, http.StatusOK, map[string]string{
		"message": "ok",
	})
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "internal server error")
	}

}
