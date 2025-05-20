package models

import "time"

type Message struct {
	ID        int64      `db:"id"`
	ChatID    int        `db:"chat_id"`
	UserID    int        `db:"user_id"`
	Text      string     `db:"text"`
	CreatedAt time.Time  `db:"created_at"`
}