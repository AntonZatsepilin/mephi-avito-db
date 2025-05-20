package models

type Listing struct {
	ID          int    `db:"id"`
	UserID      int    `db:"user_id"`
	CategoryID  int    `db:"category_id"`
	CityID	  int    `db:"city_id"`
	Title       string `db:"title"`
	Price       int    `db:"price"`
	Description string `db:"description"`
	IsActive    bool   `db:"is_active"`
	ViewCount   int    `db:"view_count"`
}