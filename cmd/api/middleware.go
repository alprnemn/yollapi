package main

import (
	"context"
	"fmt"
	cmn "github.com/alprnemn/yollapi/common"
	"net/http"
	"strconv"
	"strings"

	u "github.com/alprnemn/yollapi/cmd/api/utils"
	"github.com/golang-jwt/jwt/v5"
)

func (app *api) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if app.Config.RateLimiter.Enabled {
			if allow, retryAfter := app.RateLimiter.Allow(req.RemoteAddr); !allow {
				u.RateLimitExceededError(w, req, retryAfter.String())
				return
			}
		}
		next.ServeHTTP(w, req)
	})
}

func (app *api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			u.BadRequestResponse(w, req, cmn.ErrMissingAuthHeader)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			u.UnauthorizedError(w, req, cmn.ErrWrongAuthHeader)
			return
		}

		token, err := app.Auth.ValidateToken(parts[1])
		if err != nil {
			u.UnauthorizedError(w, req, err)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			u.UnauthorizedError(w, req, err)
			return
		}

		ctx := req.Context()

		user, err := app.Repository.User.GetByID(ctx, userID)
		if err != nil {
			u.UnauthorizedError(w, req, err)
			return
		}

		ctx = context.WithValue(ctx, cmn.UserCtx, user)
		next.ServeHTTP(w, req.WithContext(ctx))

	})
}
