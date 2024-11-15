package models

import "gorm.io/gorm"

type File struct {
	gorm.Model

	ReviewID  uint   `gorm:"not null;index"`
	Name      string `gorm:"not null"`
	MessageID uint   `gorm:"not null;index"`
}
