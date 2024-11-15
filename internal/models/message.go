package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID uint `gorm:"not null;index"`
	User   User
	ChatID uint `gorm:"not null;index"`
	Chat   Chat
	Text   string
	Files  []File
}
