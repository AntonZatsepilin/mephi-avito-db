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

func (g *GeneratorPostgres) GenerateLocation(n int) error {
    rand.Seed(time.Now().UnixNano())
    qwery := "INSERT INTO countries (name) VALUES ($1) ON CONFLICT (name) DO NOTHING"
    var country string
    var err error

    countryIDs := make([]int, 0, n)
    countryNames := make(map[string]struct{})

    for i := 0; i < n; i++ {
        country = gofakeit.Country()
        if _, exists := countryNames[country]; exists {
            continue
        }
        countryNames[country] = struct{}{}
        _, err = g.db.Exec(qwery, country)
        if err != nil {
            return err
        }
    }

    rows, err := g.db.Queryx("SELECT id FROM countries")
    if err != nil {
        return err
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            return err
        }
        countryIDs = append(countryIDs, id)
    }

    var city string
    for i := 0; i < n*100; i++ {
        city = gofakeit.City()
        countryID := countryIDs[rand.Intn(len(countryIDs))]
        _, err = g.db.Exec("INSERT INTO cities (name, country_id) VALUES ($1, $2) ON CONFLICT (country_id, name) DO NOTHING", city, countryID)
        if err != nil {
            return err
        }
    }
    return nil
}

func (g *GeneratorPostgres) GenerateCategories(n int) error {
    rand.Seed(time.Now().UnixNano())
    var err error
    for i := 0; i < n; i++ {
        category := gofakeit.ProductCategory()
        _, err := g.db.Exec("INSERT INTO categories (name) VALUES ($1)", category)
        if err != nil {
            return err
        }
    }

    type cat struct {
        ID   int    `db:"id"`
        Name string `db:"name"`
    }
    cats := []cat{}
    err = g.db.Select(&cats, "SELECT id, name FROM categories")
    if err != nil {
        return err
    }

    for i := 0; i < n; i++ {
        name := gofakeit.ProductCategory()
        parent := cats[rand.Intn(len(cats))]
        _, err := g.db.Exec("INSERT INTO categories (name, parent_id) VALUES ($1, $2)", name, parent.ID)
        if err != nil {
            return err
        }
    }
    return nil
}

