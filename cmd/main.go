package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/AntonZatsepilin/goAvitoDB.git/internal/generator"
	"github.com/AntonZatsepilin/goAvitoDB.git/internal/models"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

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
		&models.Location{},
		&models.User{},
		&models.Chat{},
		&models.Category{},
		&models.Message{},
		&models.Post{},
		&models.Password{},
		&models.Review{},
		&models.File{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	logrus.Info("database migration completed successfully")

var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 50; i++ {
		go func(i int) {
			defer wg.Done()
			if err := generator.GenerateFakeData(db, 1000); err != nil {
				logrus.Errorf("error generating fake data for iteration %d", err)
			}
		}(i)
	}
	wg.Wait()
	logrus.Info("fake data generation completed successfully")

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
