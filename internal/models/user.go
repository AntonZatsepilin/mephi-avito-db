package models

import "time"

type User struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PhoneNumber  *string   `db:"phone_number"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	ProfileImage *string   `db:"profile_image"`
	Rating       *float32  `db:"rating"`
	UserTypeID   int       `db:"user_type_id"`
	CityID       *int      `db:"city_id"`
}