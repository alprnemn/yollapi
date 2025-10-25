package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	u "github.com/alprnemn/yollapi/cmd/api/utils"
	cmn "github.com/alprnemn/yollapi/common"
	m "github.com/alprnemn/yollapi/internal/models"
	v "github.com/alprnemn/yollapi/pkg/validator"
)

// registerUserHandler godoc
//
//	@Summary		Register a new user
//	@Description	Creates a new user account with the provided information
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		models.RegisterUserPayload	true	"User registration data"
//	@Success		201		{string}	common.MessageResponse		"user created succesfully"
//	@Failure		400		{string}	string						"Bad Request"
//	@Failure		500		{string}	string						"Internal Server Error"
//	@Router			/v1/users/register [post]
func (app *api) registerUserHandler(w http.ResponseWriter, req *http.Request) {

	var payload m.RegisterUserPayload

	if err := u.ParseJSON(w, req, &payload); err != nil {
		u.BadRequestResponse(w, req, err)
		return
	}

	if err := v.ValidatePayload(payload); err != nil {
		u.BadRequestResponse(w, req, err)
		return
	}

	ctx := req.Context()

	newUser := &m.User{
		Username:  payload.Username,
		FirstName: payload.Firstname,
		LastName:  payload.Lastname,
		Email:     payload.Email,
		Phone:     payload.Phone,
		Age:       *payload.Age,
	}

	if err := newUser.Password.Set(payload.Password); err != nil {
		u.InternalServerError(w, req, err)
		return
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
//	@Success		200	{string}	[]m.User
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/v1/users [get]
func (app *api) getUsersHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	fmt.Println("who sent request: ", req.Context().Value(cmn.UserCtx).(*m.User))

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

// @Summary		login a user
// @Description	Auth a new user account with the provided information
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			payload	body		m.LoginPayload	true	"User registration data"
// @Success		200		{string}	map[string]string	"token"
// @Failure		400		{string}	string				"Bad Request"
// @Failure		404		{string}	string				"Not Found"
// @Failure		500		{string}	string				"Internal Server Error"
// @Router			/v1/auth/login [post]
func (app *api) loginHandler(w http.ResponseWriter, req *http.Request) {

	var payload m.LoginPayload

	if err := u.ParseJSON(w, req, &payload); err != nil {
		u.BadRequestResponse(w, req, err)
		return
	}

	if err := v.ValidatePayload(payload); err != nil {
		u.BadRequestResponse(w, req, err)
		return
	}

	ctx := req.Context()

	user, err := app.Repository.User.GetByUsername(ctx, payload.Username)
	if err != nil {
		u.NotFoundError(w, req, err)
		return
	}
	fmt.Println("user id: ", user.ID)

	if err := user.Password.Compare(payload.Password); err != nil {
		u.BadRequestResponse(w, req, err)
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(app.Config.JWTConfig.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.Config.JWTConfig.Issuer,
		"aud": app.Config.JWTConfig.Issuer,
	}

	token, err := app.Auth.GenerateToken(claims)
	if err != nil {
		u.InternalServerError(w, req, err)
		return
	}

	if err := u.WriteJSON(w, http.StatusOK, map[string]string{"token": token}); err != nil {
		u.InternalServerError(w, req, err)
		return
	}

}
