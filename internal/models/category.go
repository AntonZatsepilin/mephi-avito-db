package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	Listings []Listing
	Parents  []*Category `gorm:"many2many:category_relations;joinForeignKey:category_parent_id;joinReferences:category_child_id"`
	Children []*Category `gorm:"many2many:category_relations;joinForeignKey:category_child_id;joinReferences:category_parent_id"`
}
