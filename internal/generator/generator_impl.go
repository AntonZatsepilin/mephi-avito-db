package generator

import (
	"math/rand"
	"time"

	"github.com/AntonZatsepilin/goAvitoDB.git/internal/models"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

type generatorImpl struct {}

func (f *generatorImpl) Seed(seed int64) {
	gofakeit.Seed(seed)
}

func (f *generatorImpl) GenerateLocations(db *gorm.DB, count int) error {
	var locs []models.Location
	for i := 0; i < count; i++ {
		locs = append(locs, models.Location{
			City:    gofakeit.City(),
			Country: gofakeit.Country(),
		})
	}
	return db.Create(&locs).Error
}

func (f *generatorImpl) GenerateUsers(db *gorm.DB, count int) error {
	var users []models.User
	var locCount int64
	db.Model(&models.Location{}).Count(&locCount)
	for i := 0; i < count; i++ {
		users = append(users, models.User{
			Username:    gofakeit.Username(),
			Email:       gofakeit.Email(),
			PhoneNumber: gofakeit.Phone(),
			Rating:      gofakeit.Float64Range(1.0, 5.0),
			LocationID:  uint(rand.Intn(int(locCount)) + 1),
		})
	}
	return db.Create(&users).Error
}

func (f *generatorImpl) GeneratePasswords(db *gorm.DB) error {
	var users []models.User
	db.Find(&users)
	for _, u := range users {
		p := models.Password{
			UserID:    u.ID,
			Hash:      []byte(gofakeit.Password(true, true, true, true, false, 16)),
			Salt:      []byte(gofakeit.Password(true, true, true, true, false, 8)),
			Algorithm: "bcrypt",
		}
		if err := db.Create(&p).Error; err != nil {
			return err
		}
	}
	return nil
}

func (f *generatorImpl) GenerateCategories(db *gorm.DB, parents, childrenPerParent int) error {
	for i := 0; i < parents; i++ {
		parent := models.Category{Name: gofakeit.ProductCategory()}
		db.Create(&parent)
		for j := 0; j < childrenPerParent; j++ {
			child := models.Category{Name: gofakeit.ProductCategory()}
			db.Create(&child)
			if err := db.Model(&parent).Association("Children").Append(&child); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *generatorImpl) GeneratePosts(db *gorm.DB, count int) error {
	var users []models.User; db.Find(&users)
	var locs []models.Location; db.Find(&locs)
	var cats []models.Category; db.Find(&cats)
	var posts []models.Post
	for i := 0; i < count; i++ {
		posts = append(posts, models.Post{
			UserID:     users[rand.Intn(len(users))].ID,
			LocationID: locs[rand.Intn(len(locs))].ID,
			CategoryID: cats[rand.Intn(len(cats))].ID,
			Title:      gofakeit.Sentence(5),
			Description:gofakeit.Paragraph(1,3,10," "),
			Price:      gofakeit.Price(10,1000),
			IsActive:   gofakeit.Bool(),
			ViewCount:  rand.Intn(1000),
			Url:        gofakeit.URL(),
		})
	}
	return db.Create(&posts).Error
}

func (f *generatorImpl) GenerateChats(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		db.Create(&models.Chat{})
	}
	return nil
}

func (f *generatorImpl) GenerateMessages(db *gorm.DB, count int) error {
	var users []models.User; db.Find(&users)
	var chats []models.Chat; db.Find(&chats)
	var msgs []models.Message
	for i := 0; i < count; i++ {
		msgs = append(msgs, models.Message{
			UserID: users[rand.Intn(len(users))].ID,
			ChatID: chats[rand.Intn(len(chats))].ID,
			Text:   gofakeit.Sentence(10),
		})
	}
	return db.Create(&msgs).Error
}

func (f *generatorImpl) GenerateReviews(db *gorm.DB, count int) error {
	var users []models.User; db.Find(&users)
	var posts []models.Post; db.Find(&posts)
	var revs []models.Review
	for i := 0; i < count; i++ {
		revs = append(revs, models.Review{
			UserID: users[rand.Intn(len(users))].ID,
			PostID: posts[rand.Intn(len(posts))].ID,
			Comment:gofakeit.Sentence(15),
			Rating: gofakeit.Float64Range(1.0,5.0),
		})
	}
	return db.Create(&revs).Error
}

func (f *generatorImpl) GenerateUserChats(db *gorm.DB) error {
	var users []models.User; db.Find(&users)
	var chats []models.Chat; db.Find(&chats)
	for _, c := range chats {
		u1 := users[rand.Intn(len(users))].ID
		u2 := users[rand.Intn(len(users))].ID
		if u1 == u2 { continue }
		db.Model(&c).Association("Users").Append(&models.User{Model: gorm.Model{ID: u1}}, &models.User{Model: gorm.Model{ID: u2}})
	}
	return nil
}

func (f *generatorImpl) GenerateFiles(db *gorm.DB, count int) error {
	var revs []models.Review; db.Find(&revs)
	var msgs []models.Message; db.Find(&msgs)
	var files []models.File
	for i := 0; i < count; i++ {
		files = append(files, models.File{
			ReviewID:  revs[rand.Intn(len(revs))].ID,
			MessageID: msgs[rand.Intn(len(msgs))].ID,
			Name:      gofakeit.Adjective(),
			Url:       gofakeit.URL(),
		})
	}
	return db.Create(&files).Error
}

type FullGenerator struct {
	Generator
}

func (fg *FullGenerator) GenerateFakeData(db *gorm.DB, userCount int) error {
	seed := time.Now().UnixNano()
	fg.Seed(seed)
	if err := fg.GenerateLocations(db, 1000); err != nil { return err }
	if err := fg.GenerateUsers(db, userCount); err != nil { return err }
	if err := fg.GeneratePasswords(db); err != nil { return err }
	if err := fg.GenerateCategories(db, 10, 1); err != nil { return err }
	if err := fg.GeneratePosts(db, userCount/2); err != nil { return err }
	if err := fg.GenerateChats(db, userCount/2); err != nil { return err }
	if err := fg.GenerateMessages(db, userCount*2); err != nil { return err }
	if err := fg.GenerateReviews(db, userCount); err != nil { return err }
	if err := fg.GenerateUserChats(db); err != nil { return err }
	if err := fg.GenerateFiles(db, userCount/2); err != nil { return err }
	return nil
}
