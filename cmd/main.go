package main

import (
	"app"
	"app/pkg/handler"
	"app/pkg/repository"
	"app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("don't setup endv %s", err.Error())
	}

	db, err := repository.NewPostgresDB(
		repository.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_DBNAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	)

	if err != nil {
		log.Fatalf("Error while runing DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(app.Server)
	if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runing server: %s", err.Error())
	}
}
