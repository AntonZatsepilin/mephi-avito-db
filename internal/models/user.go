package models

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	Password string `db:"password_hash"`
	ProfileImage string `db:"profile_image"`
	CityID int `db:"city_id"`
	UserTypeID int `db:"user_type_id"`
}
