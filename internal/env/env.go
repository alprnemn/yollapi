package env

import (
	"github.com/lpernett/godotenv"
	"log"
)

type Config struct {
	Address string
	Env     string
	CORS    CorsConfig
}

type CorsConfig struct {
	AllowedOrigin string
}

var Envs = initConfig()

func initConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error occurred while getting envs")
	}

	return Config{
		Address: GetString("ADDRESS", ":8080"),
		Env:     "deployment",
		CORS: CorsConfig{
			AllowedOrigin: GetString("CORS_ALLOWED_ORIGIN", "http://127.0.0.1:3000"),
		},
	}

}
