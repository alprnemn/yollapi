package main

import (
	database "github.com/alprnemn/yollapi/internal/db"
	"github.com/alprnemn/yollapi/internal/env"
	"log"
)

func main() {
	app := &api{
		Config: env.Envs,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

func init() {
	db, err := database.New(
		env.Envs.Address,
	)
}
