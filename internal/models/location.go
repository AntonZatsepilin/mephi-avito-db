package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	City     string
	Country  string
	Region   string
	Users    []User
	Listings []Listing
}
