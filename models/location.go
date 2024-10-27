package models

type Location struct {
	ID      uint   `gorm:"primaryKey"`
	City    string `gorm:"size:100"`
	Region  string `gorm:"size:100"`
	Country string `gorm:"size:100"`

	Users    []User
	Listings []ListingLocation `gorm:"foreignKey:LocationID"`
}
