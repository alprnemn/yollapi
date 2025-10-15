package main

import (
	"errors"
	"fmt"
	u "github.com/alprnemn/yollapi/cmd/api/utils"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
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
			u.BadRequestResponse(w, req, errors.New("where is the auth header"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			u.UnauthorizedError(w, req, errors.New("wrong header value"))
			return
		}

		token, err := app.Auth.ValidateToken(parts[1])
		if err != nil {
			u.UnauthorizedError(w, req, err)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		fmt.Println("claims: ", claims)
		next.ServeHTTP(w, req)
	})
}
