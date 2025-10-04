package main

import (
	u "github.com/alprnemn/yollapi/cmd/api/utils"
	"net/http"
)

func (app *api) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.Config.RateLimiter.Enabled {
			if allow, retryAfter := app.RateLimiter.Allow(r.RemoteAddr); !allow {
				u.RateLimitExceededError(w, r, retryAfter.String())
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
