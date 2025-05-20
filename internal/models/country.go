package models

type Country struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}