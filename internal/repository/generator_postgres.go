package repository

import (
	"errors"
	"math/rand"
	"time"

	"github.com/AntonZatsepilin/mephi-avito-db/internal/models"
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

func (g *GeneratorPostgres) GenerateUsers(n int) error {
    rand.Seed(time.Now().UnixNano())

    // Получаем существующие user_types
    var userTypes []int
    if err := g.db.Select(&userTypes, "SELECT id FROM user_types"); err != nil {
        return err
    }
    if len(userTypes) == 0 {
        return errors.New("user types not found")
    }

    // Получаем существующие city_ids
    var cityIDs []int
    if err := g.db.Select(&cityIDs, "SELECT id FROM cities"); err != nil {
        return err
    }

    usedUsernames := make(map[string]struct{})
    usedEmails := make(map[string]struct{})

    tx, err := g.db.Beginx()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    const query = `
        INSERT INTO users (
            username, email, phone_number, password_hash, 
            created_at, profile_image, rating, user_type_id, city_id
        ) VALUES (
            :username, :email, :phone_number, :password_hash, 
            :created_at, :profile_image, :rating, :user_type_id, :city_id
        )`

    for i := 0; i < n; i++ {
        // Генерация уникального username
        username := gofakeit.Username()
        for {
            if _, exists := usedUsernames[username]; !exists {
                usedUsernames[username] = struct{}{}
                break
            }
            username = gofakeit.Username()
        }

        // Генерация уникального email
        email := gofakeit.Email()
        for {
            if _, exists := usedEmails[email]; !exists {
                usedEmails[email] = struct{}{}
                break
            }
            email = gofakeit.Email()
        }

        // Генерация остальных полей
        user := models.User{
            Username:     username,
            Email:        email,
            PhoneNumber:  gofakeit.Phone(),
            PasswordHash: "$2a$10$3YBrvN8IX/ZjWIEac5.Oxu4xGXg3Q7FmHGYcCjkrGjTZ9jML7qD4a", // Пример хэша
            CreatedAt:    randomTime(),
            ProfileImage: func() *string { s := gofakeit.URL(); return &s }(),
            Rating:       generateNullableRating(0.7),
            UserTypeID:   userTypes[rand.Intn(len(userTypes))],
            CityID:       generateNullableCityID(cityIDs, 0.8),
        }

        if _, err := tx.NamedExec(query, &user); err != nil {
            return err
        }
    }

    return tx.Commit()
}

// Вспомогательные функции
func generateNullableString(nullChance float32, generator func() string) *string {
    if rand.Float32() > nullChance {
        return nil
    }
    s := generator()
    return &s
}

func generateNullableRating(nullChance float32) *float32 {
    if rand.Float32() > nullChance {
        return nil
    }
    r := float32(rand.Intn(501)) / 100 // 0.00 - 5.00
    return &r
}

func generateNullableCityID(cityIDs []int, nullChance float32) *int {
    if len(cityIDs) == 0 || rand.Float32() < nullChance {
        return nil
    }
    id := cityIDs[rand.Intn(len(cityIDs))]
    return &id
}

func (g *GeneratorPostgres) GenerateListings(n int) error {
    rand.Seed(time.Now().UnixNano())
    // Получаем существующие ID
    var userIDs, categoryIDs, cityIDs []int
    if err := g.db.Select(&userIDs, "SELECT id FROM users"); err != nil {
        return err
    }
    if err := g.db.Select(&categoryIDs, "SELECT id FROM categories"); err != nil {
        return err
    }
    if err := g.db.Select(&cityIDs, "SELECT id FROM cities"); err != nil {
        return err
    }

    if len(userIDs) == 0 || len(categoryIDs) == 0 || len(cityIDs) == 0 {
        return errors.New("недостаточно данных в родительских таблицах")
    }

    tx, err := g.db.Beginx()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    for i := 0; i < n; i++ {
        listing := models.Listing{
            UserID:     userIDs[rand.Intn(len(userIDs))],
            CategoryID: categoryIDs[rand.Intn(len(categoryIDs))],
            CityID:     cityIDs[rand.Intn(len(cityIDs))],
            Title:      gofakeit.Sentence(5),
            Description: generateNullableString(0.3, func() string {
                return gofakeit.Paragraph(2, 3, 5, "\n")
            }),
            Price:      rand.Intn(100000) + 100,
            CreatedAt:  randomTime(),
            IsActive:   rand.Float32() < 0.8,
            ViewCount:  rand.Intn(10000),
        }

        _, err := tx.NamedExec(`
            INSERT INTO listings (
                user_id, category_id, city_id, title, 
                description, price, created_at, is_active, view_count
            ) VALUES (
                :user_id, :category_id, :city_id, :title, 
                :description, :price, :created_at, :is_active, :view_count
            )`, &listing)
        if err != nil {
            return err
        }
    }
    return tx.Commit()
}

func (g *GeneratorPostgres) GenerateReviews(n int) error {
    rand.Seed(time.Now().UnixNano())
    var userIDs []int
    var listingIDs []int
    if err := g.db.Select(&userIDs, "SELECT id FROM users"); err != nil {
        return err
    }
    if err := g.db.Select(&listingIDs, "SELECT id FROM listings WHERE is_active = true"); err != nil {
        return err
    }

    tx, err := g.db.Beginx()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    for i := 0; i < n; i++ {
        review := models.Review{
            UserID:    generateNullableID(userIDs, 0.1), // 10% chance NULL
            ListingID: listingIDs[rand.Intn(len(listingIDs))],
            Comment:   gofakeit.Sentence(10),
            Rating:    int16(rand.Intn(5) + 1), // 1-5
            CreatedAt: randomTime(),
        }

        _, err := tx.NamedExec(`
            INSERT INTO reviews (
                user_id, listing_id, comment, rating, created_at
            ) VALUES (
                :user_id, :listing_id, :comment, :rating, :created_at
            )`, &review)
        if err != nil {
            return err
        }
    }
    return tx.Commit()
}

func (g *GeneratorPostgres) GenerateChatsAndMembers(n int) error {
    rand.Seed(time.Now().UnixNano())
    var userIDs []int
    if err := g.db.Select(&userIDs, "SELECT id FROM users"); err != nil {
        return err
    }

    for i := 0; i < n; i++ {
        // Создаём чат
        var chatID int
        err := g.db.QueryRowx(`
            INSERT INTO chats (name, created_at) 
            VALUES ($1, $2) 
            RETURNING id`,
            gofakeit.Word,
            randomTime(),
        ).Scan(&chatID)
        if err != nil {
            return err
        }

        // Добавляем участников (2-10 пользователей)
        membersCount := 2
        selectedUsers := make(map[int]struct{})
        for j := 0; j < membersCount; j++ {
            userID := userIDs[rand.Intn(len(userIDs))]
            if _, exists := selectedUsers[userID]; exists {
                continue
            }
            selectedUsers[userID] = struct{}{}
            
            _, err := g.db.Exec(`
                INSERT INTO chat_members (chat_id, user_id)
                VALUES ($1, $2)
                ON CONFLICT DO NOTHING`,
                chatID,
                userID,
            )
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func (g *GeneratorPostgres) GenerateMessages(n int) error {
    rand.Seed(time.Now().UnixNano())
    var chatIDs []int
    var userIDs []int
    if err := g.db.Select(&chatIDs, "SELECT id FROM chats"); err != nil {
        return err
    }
    if err := g.db.Select(&userIDs, "SELECT id FROM users"); err != nil {
        return err
    }

    tx, err := g.db.Beginx()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    for i := 0; i < n; i++ {
        msg := models.Message{
            ChatID:    chatIDs[rand.Intn(len(chatIDs))],
            UserID:    generateNullableID(userIDs, 0.15),
            Text:      gofakeit.Sentence(rand.Intn(20) + 5),
            CreatedAt: randomTime(),
        }

        _, err := tx.NamedExec(`
            INSERT INTO messages (chat_id, user_id, text, created_at)
            VALUES (:chat_id, :user_id, :text, :created_at)`,
            &msg,
        )
        if err != nil {
            return err
        }
    }
    return tx.Commit()
}

func (g *GeneratorPostgres) GenerateFiles(n int) error {
    var messageIDs []int64
    var reviewIDs []int64
    if err := g.db.Select(&messageIDs, "SELECT id FROM messages"); err != nil {
        return err
    }
    if err := g.db.Select(&reviewIDs, "SELECT id FROM reviews"); err != nil {
        return err
    }

    tx, err := g.db.Beginx()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    for i := 0; i < n; i++ {
        file := models.File{
            Name:     gofakeit.InputName(),
            FileURL:  gofakeit.URL(),
        }

        // Выбираем случайно к чему привязать файл
        if rand.Float32() < 0.7 && len(messageIDs) > 0 { // 70% к сообщениям
            msgID := messageIDs[rand.Intn(len(messageIDs))]
            file.MessageID = &msgID
        } else if len(reviewIDs) > 0 {
            revID := reviewIDs[rand.Intn(len(reviewIDs))]
            file.ReviewID = &revID
        }

        _, err := tx.NamedExec(`
            INSERT INTO files (name, file_url, message_id, review_id)
            VALUES (:name, :file_url, :message_id, :review_id)`,
            &file,
        )
        if err != nil {
            return err
        }
    }
    return tx.Commit()
}

// Вспомогательная функция для генерации nullable ID
func generateNullableID(ids []int, nullChance float32) *int {
    if len(ids) == 0 || rand.Float32() < nullChance {
        return nil
    }
    id := ids[rand.Intn(len(ids))]
    return &id
}