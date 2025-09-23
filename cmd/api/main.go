package main

import (
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
