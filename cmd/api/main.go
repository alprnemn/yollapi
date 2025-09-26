package main

import (
	"github.com/alprnemn/yollapi/internal/config"
	database "github.com/alprnemn/yollapi/internal/db"
	"github.com/alprnemn/yollapi/internal/repository"
	"github.com/alprnemn/yollapi/internal/service"
	"log"
)

func main() {
	db, err := database.New(
		config.Envs.DbConfig.Address,
		config.Envs.DbConfig.MaxOpenConns,
		config.Envs.DbConfig.MaxIdleConns,
		config.Envs.DbConfig.MaxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	services := service.NewService(repo)

	app := &api{
		Config:     config.Envs,
		Repository: repo,
		Service:    services,
		Db:         db,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
