package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    uint `gorm:"not null;index"`
	PostID uint `gorm:"not null;index"`
	Comment   string
	Rating    float64
	Files     []File
}
