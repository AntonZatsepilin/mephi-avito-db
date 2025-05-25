package main

import (
	"os"
	"sync"

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
// 1. Параллельная генерация независимых сущностей
var wg sync.WaitGroup
wg.Add(2)

go func() {
    defer wg.Done()
    repos.Generator.GenerateLocation(200)
    logrus.Info("locations generated")
}()

go func() {
    defer wg.Done()
    repos.Generator.GenerateCategories(1000)
    logrus.Info("categories generated")
}()

wg.Wait()

// 2. Генерация пользователей с внутренним параллелизмом
repos.Generator.GenerateUsers(50000)
logrus.Info("users generated")

// 3. Параллельная генерация зависимых сущностей
wg.Add(4)

// Листинги
go func() {
    defer wg.Done()
    repos.Generator.GenerateListings(7000)
    logrus.Info("listings generated")
}()

// Чаты
go func() {
    defer wg.Done()
    repos.Generator.GenerateChatsAndMembers(10000)
    logrus.Info("chats and members generated")
}()

// Файлы
go func() {
    defer wg.Done()
    repos.Generator.GenerateFiles(20000)
    logrus.Info("files generated")
}()

// Отзывы и сообщения последовательно
go func() {
    defer wg.Done()
    repos.Generator.GenerateReviews(100000)
    logrus.Info("reviews generated")
    repos.Generator.GenerateMessages(100000)
    logrus.Info("messages generated")
}()

wg.Wait()


	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
