package generator

import "gorm.io/gorm"

// Generator defines methods to populate each table with fake data.
type Generator interface {
	Seed(seed int64)
	GenerateLocations(db *gorm.DB, count int) error
	GenerateUsers(db *gorm.DB, count int) error
	GeneratePasswords(db *gorm.DB) error
	GenerateCategories(db *gorm.DB, parents, childrenPerParent int) error
	GeneratePosts(db *gorm.DB, count int) error
	GenerateChats(db *gorm.DB, count int) error
	GenerateMessages(db *gorm.DB, count int) error
	GenerateReviews(db *gorm.DB, count int) error
	GenerateUserChats(db *gorm.DB) error
	GenerateFiles(db *gorm.DB, count int) error
}

func NewGenerator() Generator {
	return &generatorImpl{}
}