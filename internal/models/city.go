package models

type City struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	CountryID int    `db:"country_id"`
}