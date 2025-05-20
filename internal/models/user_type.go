package models

type UserType struct {
	ID   int    `db:"id"`
	Type string `db:"type"` // admin, user, moderator
}