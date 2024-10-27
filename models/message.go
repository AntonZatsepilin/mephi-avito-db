package models

type Message struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ChatID    uint
	Text      string `gorm:"type:text"`
	Timestamp string `gorm:"type:timestamp"`

	User  User
	Chat  Chat
	Files []MessageFile `gorm:"foreignKey:MessageID"`
}
