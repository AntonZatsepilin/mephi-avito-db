package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LocationID   uint
	Location     Location
	Username     string
	Email        string `gorm:"unique"`
	PhoneNumber  string
	PasswordHash string
	DateJoined   time.Time
	ProfileImage string
	Rating       float64
	Listings     []Listing
	Reviews      []Review
	Messages     []Message
	ChatsAsUser1 []Chat `gorm:"foreignKey:User1ID"`
	ChatsAsUser2 []Chat `gorm:"foreignKey:User2ID"`
}
