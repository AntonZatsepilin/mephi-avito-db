package models

import "gorm.io/gorm"

type Listing struct {
	gorm.Model
	CategoryID  uint
	Category    Category
	UserID      uint
	User        User
	LocationID  uint
	Location    Location
	Title       string
	Description string
	Price       float64
	IsActive    bool
	ViewCount   int
	URLs        string
	Reviews     []Review
}
