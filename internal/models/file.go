package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name     string
	URL      string
	Messages []MessageFile
	Reviews  []ReviewFile
}
