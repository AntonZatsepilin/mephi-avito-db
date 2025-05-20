package models

import (
	"time"
)

type Listing struct {
	ID          int64           `db:"id"`
	UserID      int             `db:"user_id"`
	CategoryID  int             `db:"category_id"`
	CityID      int             `db:"city_id"`
	Title       string          `db:"title"`
	Description *string         `db:"description"`
	Price       int `db:"price"`
	CreatedAt   time.Time       `db:"created_at"`
	IsActive    bool            `db:"is_active"`
	ViewCount   int             `db:"view_count"`
}