package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Posts    []Post
	Children []*Category `gorm:"many2many:category_children"`
}
