package models

type MessageFile struct {
	MessageID uint `gorm:"primaryKey"`
	FileID    uint `gorm:"primaryKey"`

	Message Message
	File    File
}
