package main

import (
	"database/sql"
	"errors"
	"github.com/alprnemn/yollapi/internal/auth"
	"log"
	"net/http"
	"time"

	"github.com/alprnemn/yollapi/internal/ratelimiter"

	_ "github.com/alprnemn/yollapi/docs"
	cfg "github.com/alprnemn/yollapi/internal/config"
	"github.com/alprnemn/yollapi/internal/repository"
	"github.com/alprnemn/yollapi/internal/service"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type api struct {
	Config      cfg.Config
	Repository  *repository.Repository
	Service     *service.Service
	Db          *sql.DB
	RateLimiter *ratelimiter.FixedWindowRateLimiter
	Auth        *auth.JWTAuthenticator
}

func (app *api) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(app.RateLimiterMiddleware)

	// put cors middleware before the rate limiter
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.Envs.CORS.AllowedOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// v1
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		//docsURL := "http://127.0.0.1:8080/swagger/doc.json"
		r.Handle("/swagger/*", httpSwagger.WrapHandler)

		r.Route("/users", func(r chi.Router) {
			r.Post("/register", app.registerUserHandler)

			r.Group(func(r chi.Router) {
				r.Use(app.AuthMiddleware)
				r.Get("/", app.getUsersHandler)
			})
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", app.loginHandler)
		})

	})
	return r
}

func (app *api) run(mux http.Handler) error {

	server := &http.Server{
		Addr:         app.Config.Address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at 127.0.0.1%s", app.Config.Address)
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
