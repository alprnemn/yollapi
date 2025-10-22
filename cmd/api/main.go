package main

import (
	"github.com/alprnemn/yollapi/internal/auth"
	"github.com/alprnemn/yollapi/internal/cache"
	cfg "github.com/alprnemn/yollapi/internal/config"
	database "github.com/alprnemn/yollapi/internal/db"
	"github.com/alprnemn/yollapi/internal/ratelimiter"
	"github.com/alprnemn/yollapi/internal/repository"
	"github.com/alprnemn/yollapi/internal/service"
	"log"
)

const version = "1.1.0"

// @title			Yolla Api
// @version		1.1.0
// @description	This is an Api for Yolla App
// @termsOfService	http://swagger.io/terms/
func main() {
	db, err := database.New(
		cfg.Envs.DbConfig.Address,
		cfg.Envs.DbConfig.MaxOpenConns,
		cfg.Envs.DbConfig.MaxIdleConns,
		cfg.Envs.DbConfig.MaxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rateLimiter := ratelimiter.NewFixedWindowRateLimiter(
		cfg.Envs.RateLimiter.RequestsPerTimeFrame,
		cfg.Envs.RateLimiter.TimeFrame,
	)

	repo := repository.NewRepository(db)
	services := service.NewService(repo)

	authenticator := auth.NewJWTAuthenticator(
		cfg.Envs.JWTConfig.Secret,
		cfg.Envs.JWTConfig.Issuer,
		cfg.Envs.JWTConfig.Issuer)

	redisClient := cache.NewRedisClient(
		cfg.Envs.RedisConfig.Address,
		cfg.Envs.RedisConfig.Password,
		cfg.Envs.RedisConfig.DB)

	redisRepository := cache.NewRedisRepository(redisClient)

	app := &api{
		Config:      cfg.Envs,
		Repository:  repo,
		Service:     services,
		Db:          db,
		RateLimiter: rateLimiter,
		Auth:        authenticator,
		Cache:       redisRepository,
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatal("server error:", err)
	}

}
