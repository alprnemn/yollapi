package main

import (
	"github.com/alprnemn/yollapi/internal/types"
	"net/http"
)

func (app *api) registerUserHandler(w http.ResponseWriter, req *http.Request) {

	var RegisterUserPayload types.RegisterUserPayload

	if err := ParseJSON(w, req, &RegisterUserPayload); err != nil {
		app.badRequestResponse(w, req, err)
		return
	}

	ctx := req.Context()

	newUser := &types.User{
		Username:  RegisterUserPayload.Username,
		FirstName: RegisterUserPayload.Firstname,
		LastName:  RegisterUserPayload.Lastname,
		Email:     RegisterUserPayload.Email,
		Phone:     RegisterUserPayload.Phone,
		Password:  RegisterUserPayload.Password,
		Age:       RegisterUserPayload.Age,
	}

	if err := app.Store.Users.Create(ctx, newUser); err != nil {
		app.internalServerError(w, req, err)
		return
	}

	if err := WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created"}); err != nil {
		app.internalServerError(w, req, err)
		return
	}

}
