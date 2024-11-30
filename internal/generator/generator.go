package generator

import (
	"math/rand"
	"time"

	"github.com/AntonZatsepilin/goAvitoDB.git/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func GenerateFakeData(db *gorm.DB, userCount int) error {
	gofakeit.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	// Генерация локаций
	var locations []models.Location
	for i := 0; i < 100; i++ {
		location := models.Location{
			City:    gofakeit.City(),
			Country: gofakeit.Country(),
		}
		locations = append(locations, location)
	}

	if err := db.Create(&locations).Error; err != nil {
		return err
	}

	// Генерация пользователей
	var users []models.User
	for i := 0; i < userCount; i++ {
		user := models.User{
			Username:    gofakeit.Username(),
			Email:       gofakeit.Email(),
			PhoneNumber: gofakeit.Phone(),
			Rating:      gofakeit.Float64Range(1.0, 5.0),
			LocationID:  uint(rand.Intn(10) + 1), // Ссылка на Location
		}
		users = append(users, user)
	}

	// Сохранение пользователей в базу
	if err := db.Create(&users).Error; err != nil {
		return err
	}

	// Генерация категорий
	var categories []models.Category
	for i := 0; i < 5; i++ {
		category := models.Category{
			Name: gofakeit.Word(),
		}
		categories = append(categories, category)
	}

	if err := db.Create(&categories).Error; err != nil {
		return err
	}

	// Генерация постов
	var posts []models.Post
	for i := 0; i < userCount/2; i++ {
		post := models.Post{
			UserID:      users[rand.Intn(len(users))].ID,
			LocationID:  locations[rand.Intn(len(locations))].ID,
			CategoryID:  categories[rand.Intn(len(categories))].ID,
			Title:       gofakeit.Sentence(5),
			Description: gofakeit.Paragraph(1, 3, 10, " "),
			Price:       gofakeit.Price(10, 1000),
			IsActive:    gofakeit.Bool(),
			ViewCount:   rand.Intn(1000),
			Url:         gofakeit.URL(),
		}
		posts = append(posts, post)
	}

	if err := db.Create(&posts).Error; err != nil {
		return err
	}

	// Генерация чатов и сообщений
	var chats []models.Chat
	for i := 0; i < userCount/3; i++ {
		chat := models.Chat{}
		chats = append(chats, chat)
	}

	if err := db.Create(&chats).Error; err != nil {
		return err
	}

	var messages []models.Message
	for i := 0; i < userCount*2; i++ {
		message := models.Message{
			UserID: users[rand.Intn(len(users))].ID,
			ChatID: chats[rand.Intn(len(chats))].ID,
			Text:   gofakeit.Sentence(10),
		}
		messages = append(messages, message)
	}

	if err := db.Create(&messages).Error; err != nil {
		return err
	}

	// Генерация отзывов
	var reviews []models.Review
	for i := 0; i < userCount; i++ {
		review := models.Review{
			UserID:    users[rand.Intn(len(users))].ID,
			ListingID: posts[rand.Intn(len(posts))].ID,
			Comment:   gofakeit.Sentence(15),
			Rating:    gofakeit.Float64Range(1.0, 5.0),
		}
		reviews = append(reviews, review)
	}

	if err := db.Create(&reviews).Error; err != nil {
		return err
	}

	var userIDs []uint
	var chatIDs []uint

	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	for _, chat := range chats {
		chatIDs = append(chatIDs, chat.ID)
	}

	for _, chatID := range chatIDs {
		user1ID := userIDs[rand.Intn(len(users))]
		user2ID := userIDs[rand.Intn(len(users))]
		if user1ID == user2ID {
			continue
		}
		userChats := make([]map[string]interface{}, 0, 2)
		userChats = append(userChats, map[string]interface{}{"user_id": user1ID, "chat_id": chatID})
		userChats = append(userChats, map[string]interface{}{"user_id": user2ID, "chat_id": chatID})
		if err := db.Table("user_chats").Create(&userChats).Error; err != nil {
			panic(err)
		}
	}

	// Генерация файлов
	var files []models.File
	for i := 0; i < userCount/2; i++ {
		file := models.File{
			ReviewID:  reviews[rand.Intn(len(reviews))].ID,
			Name:      gofakeit.Name(),
			MessageID: messages[rand.Intn(len(messages))].ID,
			Url:       gofakeit.URL(),
		}
		files = append(files, file)
	}

	if err := db.Create(&files).Error; err != nil {
		return err
	}

	// Генерация вложенных категорий
	for i := 0; i < userCount; i++ {
		parentCategory := models.Category{
			Name: gofakeit.Word(),
		}
		if err := db.Create(&parentCategory).Error; err != nil {
			return err
		}

		for j := 0; j < userCount; j++ {
			childCategory := models.Category{
				Name: gofakeit.Word(),
			}
			if err := db.Create(&childCategory).Error; err != nil {
				return err
			}
			// Создаем связь между категориями
			if err := db.Model(&parentCategory).Association("Children").Append(&childCategory); err != nil {
				return err
			}
		}
	}
	return nil
}
