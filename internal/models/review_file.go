package models

import "gorm.io/gorm"

type ReviewFile struct {
	gorm.Model
	ReviewID uint
	Review   Review `gorm:"foreignKey:ReviewID"`
	FileID   uint
	File     File `gorm:"foreignKey:FileID"`
}
