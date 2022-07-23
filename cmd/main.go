package main

import (
	"app/pkg"
	"app/pkg/handler"
	"app/pkg/repository"
	"app/pkg/service"
	"github.com/caarlos0/env/v6"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config := pkg.Config{}

	if err := env.Parse(&config); err != nil {
		log.Fatalf("Error reading config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config)

	if err != nil {
		log.Fatalf("Error while runing DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(pkg.Server)
	if err := srv.Run(config.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runing server: %s", err.Error())
	}
}
