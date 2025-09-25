package main

import (
	"database/sql"
	"errors"
	"github.com/alprnemn/yollapi/internal/env"
	"github.com/alprnemn/yollapi/internal/store"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"time"
)

type api struct {
	Config env.Config
	Store  *store.Storage
	db     *sql.DB
}

func (app *api) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// put cors middleware before the rate limiter
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.Envs.CORS.AllowedOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// r.Use(app.RateLimiterMiddleware)

	// v1
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/users", func(r chi.Router) {
			r.Post("/register", app.registerUserHandler)
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
	log.Printf("server has started at %s", app.Config.Address)
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
