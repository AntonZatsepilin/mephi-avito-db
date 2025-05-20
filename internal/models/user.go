package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LocationID  uint   `gorm:"not null;index"`
	Username    string `gorm:"not null;index"`
	Email       string `gorm:"unique;not null;unique"`
	PhoneNumber string `gorm:"unique;not null"`
	Rating      float64
	Password    Password
	Posts       []Post
	Reviews     []Review
	Chats       []Chat `gorm:"many2many:user_chats;"`
}
