package main

import (
	database "github.com/alprnemn/yollapi/internal/db"
	"github.com/alprnemn/yollapi/internal/env"
	"github.com/alprnemn/yollapi/internal/store"
	"log"
)

func main() {
	db, err := database.New(
		env.Envs.DbConfig.Address,
		env.Envs.DbConfig.MaxOpenConns,
		env.Envs.DbConfig.MaxIdleConns,
		env.Envs.DbConfig.MaxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := store.NewStorage(db)

	app := &api{
		Config: env.Envs,
		Store:  store,
		db:     db,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
