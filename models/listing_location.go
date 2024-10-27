package models

type ListingLocation struct {
	ListingID  uint `gorm:"primaryKey"`
	LocationID uint `gorm:"primaryKey"`

	Listing  Listing
	Location Location
}
