package config

import (
	"log"
	"time"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Address     string
	ApiURL      string
	Env         string
	DbConfig    DbConfig
	CORS        CorsConfig
	RateLimiter RLConfig
	JWTConfig   JWTConfig
}

type CorsConfig struct {
	AllowedOrigin string
}

type JWTConfig struct {
	Secret string
	Exp    time.Duration
	Issuer string
}

type DbConfig struct {
	Address      string
	Port         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type RLConfig struct {
	RequestsPerTimeFrame int
	TimeFrame            time.Duration
	Enabled              bool
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error occurred while getting envs")
	}

	return Config{
		Address: GetString("ADDRESS", ":8080"),
		ApiURL:  GetString("PUBLIC_HOST", "http://127.0.0.1"),
		Env:     "deployment",
		DbConfig: DbConfig{
			Address:      GetString("DB_ADDR", "postgres://user:adminpassword@localhost/yollapi?sslmode=disable"),
			MaxOpenConns: GetInt("DB_MAXOPENCONNS", 3),
			MaxIdleConns: GetInt("DB_MAXIDLECONNS", 3),
			MaxIdleTime:  GetString("DB_MAXIDLETIME", "15min"),
		},
		CORS: CorsConfig{
			AllowedOrigin: GetString("CORS_ALLOWED_ORIGIN", "http://127.0.0.1:3000"),
		},
		RateLimiter: RLConfig{
			RequestsPerTimeFrame: 1,
			TimeFrame:            3 * time.Second,
			Enabled:              GetBool("RATE_LIMITER_ENABLED", true),
		},
		JWTConfig: JWTConfig{
			Secret: GetString("JWT_SECRET_KEY", "dddddd"),
			Issuer: GetString("JWT_ISSUER", "ISSUER"),
			Exp:    time.Second * time.Duration(GetInt("JWT_EXP_SECOND", 15)),
		},
	}
}
