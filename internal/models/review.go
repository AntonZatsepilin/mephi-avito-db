package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID    uint
	User      User
	ListingID uint
	Comment   string
	Rating    float64
	Files     []ReviewFile
}
