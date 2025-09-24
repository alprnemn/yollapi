package main

import (
	"log"
	"net/http"
)

func (app *api) internalServerError(w http.ResponseWriter, req *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", req.Method, req.URL.Path, err.Error())
	WriteError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *api) badRequestResponse(w http.ResponseWriter, req *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", req.Method, req.URL.Path, err)
	WriteError(w, http.StatusBadRequest, err.Error())
}

func (app *api) notFoundError(w http.ResponseWriter, req *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", req.Method, req.URL.Path, err)
	WriteError(w, http.StatusNotFound, "resource not found")
}
