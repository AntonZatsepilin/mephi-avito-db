package models

import "time"

type Chat struct {
	ID        int        `db:"id"`
	Name      *string    `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
}