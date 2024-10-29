package models

import "gorm.io/gorm"

type MessageFile struct {
	gorm.Model
	MessageID uint
	Message   Message
	FileID    uint
	File      File
}
