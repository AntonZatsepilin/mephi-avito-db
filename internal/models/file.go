package models

type File struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	FileURL   string  `db:"file_url"`
	MessageID *int64  `db:"message_id"`
	ReviewID  *int64  `db:"review_id"`
}