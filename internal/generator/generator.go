package generator

// import (
// 	"time"

// 	"github.com/AntonZatsepilin/goAvitoDB/models"
// 	"github.com/jaswdr/faker/v2"
// 	"github.com/sirupsen/logrus"
// 	"gorm.io/gorm"
// )

// func GenerateFakeData(db *gorm.DB) {
// 	fake := faker.New()

// 	for i := 0; i < 2; i++ {
// 		location := models.Location{
// 			City:    fake.Address().City(),
// 			Country: fake.Address().Country(),
// 		}
// 		db.Create(&location)

// 		username := fake.Person().Name()
// 		if len(username) > 20 {
// 			username = username[:20]
// 		}

// 		user := models.User{
// 			Username:     username,
// 			Email:        fake.Internet().Email(),
// 			PhoneNumber:  fake.Phone().Number(),
// 			PasswordHash: fake.Internet().Password(),
// 			DateJoined:   fake.Time().Time(time.Now().AddDate(-1, 0, 0)).Format(time.RFC3339),
// 			ProfileImage: fake.Internet().URL(),
// 			Rating:       fake.Float64(1, 0, 5),
// 			LocationID:   location.ID,
// 		}
// 		db.Create(&user)

// 		category := models.Category{
// 			Name: fake.Lorem().Word(),
// 		}
// 		db.Create(&category)

// 		listing := models.Listing{
// 			Title:       fake.Lorem().Word(),
// 			Description: fake.Lorem().Sentence(10),
// 			Price:       fake.Float64(2, 10, 1000),
// 			UserID:      user.ID,
// 			CategoryID:  category.ID,
// 			DateCreated: fake.Time().Time(time.Now().AddDate(-1, 0, 0)).Format(time.RFC3339),
// 			IsActive:    true,
// 			ViewCount:   uint(fake.IntBetween(0, 1000)),
// 		}
// 		db.Create(&listing)

// 		listingLocation := models.ListingLocation{
// 			ListingID:  listing.ID,
// 			LocationID: location.ID,
// 		}
// 		db.Create(&listingLocation)

// 		chat := models.Chat{
// 			CreatedAt: fake.Time().Time(time.Now().AddDate(-1, 0, 0)).Format(time.RFC3339),
// 		}
// 		db.Create(&chat)

// 		message := models.Message{
// 			Text:      fake.Lorem().Sentence(20),
// 			UserID:    user.ID,
// 			ChatID:    chat.ID,
// 			Timestamp: fake.Time().Time(time.Now().AddDate(-1, 0, 0)).Format(time.RFC3339),
// 		}
// 		db.Create(&message)

// 		review := models.Review{
// 			Comment:    fake.Lorem().Sentence(15),
// 			UserID:     user.ID,
// 			ListingID:  listing.ID,
// 			DatePosted: fake.Time().Time(time.Now().AddDate(-1, 0, 0)).Format(time.RFC3339),
// 			Rating:     fake.Float64(1, 0, 5),
// 		}
// 		db.Create(&review)
// 	}

// 	logrus.Info("fake data generation completed successfully")
// }
