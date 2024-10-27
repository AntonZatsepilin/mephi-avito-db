package models

type Listing struct {
	ID          uint `gorm:"primaryKey"`
	CategoryID  uint
	UserID      uint
	Title       string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Price       float64
	DateCreated string `gorm:"type:timestamp"`
	IsActive    bool
	ViewCount   uint
	URLS        string `gorm:"type:text"`

	User     User
	Category Category
	Location []ListingLocation `gorm:"foreignKey:ListingID"`
	Reviews  []Review
}
