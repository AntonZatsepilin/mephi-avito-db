package repository

import (
	"github.com/jmoiron/sqlx"
)

type Generator interface {
	GenerateCountries(n int) error
}

type Repository struct {
Generator
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Generator: NewGeneratorPostgres(db),
	}
}
