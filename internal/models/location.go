package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Users   []User
	Posts   []Post
	City    string `gorm:"not null"`
	Country string `gorm:"not null"`
}
