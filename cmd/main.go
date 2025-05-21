package main

import (
	"os"

	"github.com/AntonZatsepilin/mephi-avito-db/internal/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBname:   os.Getenv("DB_NAME"),
		SSLmode:  os.Getenv("DB_SSL_MODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	
	logrus.Info("app initialized")

	repos.Generator.GenerateLocation(10)
	logrus.Info("locations generated")
	repos.Generator.GenerateCategories(10)
	logrus.Info("categories generated")
	repos.Generator.GenerateUsers(100)
	logrus.Info("users generated")


	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
