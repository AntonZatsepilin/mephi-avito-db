package repository

import (
	"github.com/jmoiron/sqlx"
)

type Generator interface {
	GenerateLocation(n int) error
	GenerateCategories(n int) error
	GenerateUsers(n int) error
	GenerateListings(n int) error
	GenerateReviews(n int) error
	GenerateChatsAndMembers(n int) error
	GenerateMessages(n int) error
	GenerateFiles(n int) error
}

type Repository struct {
Generator
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Generator: NewGeneratorPostgres(db),
	}
}
