package repository

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
)

type GeneratorPostgres struct {
	db *sqlx.DB
}

func NewGeneratorPostgres(db *sqlx.DB) *GeneratorPostgres {
	return &GeneratorPostgres{db: db}
}

func randomTime() time.Time {
    return time.Now().Add(-time.Duration(rand.Intn(8760)) * time.Hour) 
}

func randomBool() bool {
    return rand.Intn(2) == 1
}

func (g *GeneratorPostgres) GenerateCountries(n int) error {
	qwery := "INSERT INTO countries (name) VALUES ($1) ON CONFLICT (name) DO NOTHING"
	var country string

var err error
for i := 0; i < n; i++ {
		country = gofakeit.Country()
		_, err = g.db.Exec(qwery, country)
		if err != nil {
			return err
		}
}
return nil
}
