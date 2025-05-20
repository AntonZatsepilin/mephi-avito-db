package repository

import "github.com/jmoiron/sqlx"

type GeneratorPostgres struct {
	db *sqlx.DB
}

func NewGeneratorPostgres(db *sqlx.DB) *GeneratorPostgres {
	return &GeneratorPostgres{db: db}
}

