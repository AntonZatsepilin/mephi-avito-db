package models

type ReviewFile struct {
	ReviewID uint `gorm:"primaryKey"`
	FileID   uint `gorm:"primaryKey"`

	Review Review
	File   File
}
