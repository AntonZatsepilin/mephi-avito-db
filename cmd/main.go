package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AntonZatsepilin/goAvitoDB/models"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("db.port"),
		viper.GetString("db.sslmode"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Chat{},
		&models.ChatUser{},
		&models.Category{},
		&models.CategoryRelation{},
		&models.Listing{},
		&models.ListingLocation{},
		&models.Location{},
		&models.Message{},
		&models.MessageFile{},
		&models.Review{},
		&models.ReviewFile{},
		&models.File{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	logrus.Info("database migration completed successfully")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}
