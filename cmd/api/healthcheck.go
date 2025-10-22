package main

import (
	"fmt"
	"github.com/alprnemn/yollapi/cmd/api/utils"
	"net/http"
)

func (app *api) healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("health")
	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "ok",
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
	}

}
