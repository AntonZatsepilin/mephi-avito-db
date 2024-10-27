package models

type Chat struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt string     `gorm:"type:timestamp"`
	Users     []ChatUser `gorm:"foreignKey:ChatID"`
	Messages  []Message
}
