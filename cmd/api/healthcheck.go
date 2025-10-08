package main

import (
	"fmt"
	"net/http"

	"github.com/alprnemn/yollapi/cmd/api/utils"
)

func (app *api) healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("health")
	err := utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "ok",
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
	}

}
