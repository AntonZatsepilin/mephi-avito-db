package models

type File struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"size:255"`
	FileURL string `gorm:"size:255"`

	MessageFiles []MessageFile
	ReviewFiles  []ReviewFile
}
