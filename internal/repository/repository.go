package repository

import (
	"github.com/jmoiron/sqlx"
)

type Generator interface {
	GenerateLocation(n int) error
	GenerateCategories(n int) error
	GenerateUsers(n int) error
}

type Repository struct {
Generator
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Generator: NewGeneratorPostgres(db),
	}
}
