package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID      uint `gorm:"not null;index"`
	LocationID  uint `gorm:"not null;index"`
	CategoryID  uint `gorm:"not null;index"`
	Title       string
	Description string
	Price       float64
	IsActive    bool
	ViewCount   int
	Url         string
	Reviews []Review `gorm:"foreignKey:PostID"`

}
