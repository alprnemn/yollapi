package main

import (
	"context"
	"fmt"
	cmn "github.com/alprnemn/yollapi/common"
	m "github.com/alprnemn/yollapi/internal/models"
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

		// check auth header
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			u.BadRequestResponse(w, req, cmn.ErrMissingAuthHeader)
			return
		}

		// split header into 2 parts like ['Bearer','eqe6t8q67e8tq.3df24f24f24.as4fsa4fas4f4sa']
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			u.UnauthorizedError(w, req, cmn.ErrWrongAuthHeader)
			return
		}

		// valide token that come from user
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

		user, err := app.getUser(ctx, userID)
		if err != nil {
			u.UnauthorizedError(w, req, err)
			return
		}

		ctx = context.WithValue(ctx, cmn.UserCtx, user)
		next.ServeHTTP(w, req.WithContext(ctx))

	})
}

func (app *api) getUser(ctx context.Context, userID int64) (*m.User, error) {

	redisEnabled := app.Config.RedisConfig.Enabled

	if !redisEnabled {
		return app.Repository.User.GetByID(ctx, userID)
	}

	user, _ := app.Cache.User.Get(ctx, userID)

	if user == nil {
		user, err := app.Repository.User.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		if err := app.Cache.User.Set(ctx, user); err != nil {
			return nil, err
		}
	}

	return user, nil

}
