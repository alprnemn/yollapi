package env

import (
	"github.com/lpernett/godotenv"
	"log"
)

type Config struct {
	Address  string
	Env      string
	DbConfig DbConfig
	CORS     CorsConfig
}

type CorsConfig struct {
	AllowedOrigin string
}

type DbConfig struct {
	Address      string
	Port         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
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
		DbConfig: DbConfig{
			Address:      GetString("DB_ADDR", "postgres://user:adminpassword@localhost/yollapi?sslmode=disable"),
			MaxOpenConns: GetInt("DB_MAXOPENCONNS", 3),
			MaxIdleConns: GetInt("DB_MAXIDLECONNS", 3),
			MaxIdleTime:  GetString("DB_MAXIDLETIME", "15min"),
		},
		CORS: CorsConfig{
			AllowedOrigin: GetString("CORS_ALLOWED_ORIGIN", "http://127.0.0.1:3000"),
		},
	}

}
