package main

import (
	u "github.com/alprnemn/yollapi/cmd/api/utils"
	cmn "github.com/alprnemn/yollapi/common"
	"github.com/alprnemn/yollapi/internal/domain"
	v "github.com/alprnemn/yollapi/pkg/validator"
	"net/http"
)

// registerUserHandler godoc
//
//	@Summary		Register a new user
//	@Description	Creates a new user account with the provided information
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		domain.RegisterUserPayload	true	"User registration data"
//	@Success		201		{string}	common.MessageResponse		"user created succesfully"
//	@Failure		400		{string}	string						"Bad Request"
//	@Failure		500		{string}	string						"Internal Server Error"
//	@Router			/v1/users/register [post]
func (app *api) registerUserHandler(w http.ResponseWriter, req *http.Request) {

	var payload domain.RegisterUserPayload

	if err := u.ParseJSON(w, req, &payload); err != nil {
		u.BadRequestResponse(w, req, err)
		return
	}

	if err := v.ValidatePayload(payload); err != nil {
		u.BadRequestResponse(w, req, err)
		return
	}

	ctx := req.Context()

	newUser := &domain.User{
		Username:  payload.Username,
		FirstName: payload.Firstname,
		LastName:  payload.Lastname,
		Email:     payload.Email,
		Phone:     payload.Phone,
		Password:  payload.Password,
		Age:       payload.Age,
	}

	if err := app.Service.User.Register(ctx, newUser); err != nil {
		if err.Code == http.StatusConflict {
			u.ConflictError(w, req, err)
			return
		}
		u.DatabaseError(w, req, err)
		return
	}

	if err := u.WriteJSON(w, http.StatusCreated, cmn.MessageResponse{
		Message: "user created successfully",
	}); err != nil {
		u.InternalServerError(w, req, err)
		return
	}

}

// getUsersHandler godoc
//
//	@Summary		Get all users
//	@Description	Gets all users from db
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	[]domain.User
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/v1/users [get]
func (app *api) getUsersHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	users, err := app.Service.User.GetAllUsers(ctx)
	if err != nil {
		u.DatabaseError(w, req, err)
		return
	}

	if err := u.WriteJSON(w, http.StatusOK, users); err != nil {
		u.InternalServerError(w, req, err)
		return
	}

}
