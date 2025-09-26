package main

import (
	"github.com/alprnemn/yollapi/internal/domain"
	"log"
	"net/http"
)

func (app *api) registerUserHandler(w http.ResponseWriter, req *http.Request) {
	log.Print("from handler")
	var RegisterUserPayload domain.RegisterUserPayload

	if err := ParseJSON(w, req, &RegisterUserPayload); err != nil {
		app.badRequestResponse(w, req, err)
		return
	}

	ctx := req.Context()

	newUser := &domain.User{
		Username:  RegisterUserPayload.Username,
		FirstName: RegisterUserPayload.Firstname,
		LastName:  RegisterUserPayload.Lastname,
		Email:     RegisterUserPayload.Email,
		Phone:     RegisterUserPayload.Phone,
		Password:  RegisterUserPayload.Password,
		Age:       RegisterUserPayload.Age,
	}

	if err := app.Service.User.Register(ctx, newUser); err != nil {
		app.internalServerError(w, req, err)
		return
	}

	if err := WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created"}); err != nil {
		app.internalServerError(w, req, err)
		return
	}

}
