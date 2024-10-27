package models

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex;size:50"`
	Email        string `gorm:"uniqueIndex;size:100"`
	PhoneNumber  string `gorm:"size:20"`
	PasswordHash string `gorm:"size:255"`
	DateJoined   string `gorm:"type:timestamp"`
	ProfileImage string `gorm:"size:255"`
	Rating       float64

	LocationID uint
	Location   Location

	Listings []Listing
	Reviews  []Review
	Chats    []ChatUser `gorm:"foreignKey:UserID"`
	Messages []Message  `gorm:"foreignKey:UserID"`
}
