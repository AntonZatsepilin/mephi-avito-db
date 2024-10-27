package models

type Review struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	ListingID  uint
	Comment    string `gorm:"type:text"`
	DatePosted string `gorm:"type:timestamp"`
	Rating     float64

	User    User
	Listing Listing
	Files   []ReviewFile `gorm:"foreignKey:ReviewID"`
}
