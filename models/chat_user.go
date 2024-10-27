package models

type ChatUser struct {
	ChatID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`

	Chat Chat
	User User
}
