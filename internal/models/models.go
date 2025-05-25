package models

import (
	"time"
)

type UserType struct {
    ID   int    `db:"id"`
    Type string `db:"type"`
}

type Country struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
}

type City struct {
    ID        int    `db:"id"`
    Name      string `db:"name"`
    CountryID int    `db:"country_id"`
}

type Category struct {
    ID       int    `db:"id"`
    Name     string `db:"name"`
    ParentID *int   `db:"parent_id"`
}

type User struct {
    ID            int        `db:"id"`
    Username      string     `db:"username"`
    Email         string     `db:"email"`
    PhoneNumber   string    `db:"phone_number"`
    PasswordHash  string     `db:"password_hash"`
    CreatedAt     time.Time  `db:"created_at"`
    ProfileImage  *string    `db:"profile_image"`
    Rating        *float32   `db:"rating"`
    UserTypeID    int        `db:"user_type_id"`
    CityID        *int       `db:"city_id"`
}

type Listing struct {
    ID          int64           `db:"id"`
    UserID      int             `db:"user_id"`
    CategoryID  int             `db:"category_id"`
    CityID      int             `db:"city_id"`
    Title       string          `db:"title"`
    Description *string         `db:"description"`
    Price       int `db:"price"`
    CreatedAt   time.Time       `db:"created_at"`
    IsActive    bool            `db:"is_active"`
    ViewCount   int             `db:"view_count"`
}

type Review struct {
    ID        int64      `db:"id"`
    UserID    *int       `db:"user_id"`
    ListingID int        `db:"listing_id"`
    Comment   string    `db:"comment"`
    Rating    int16      `db:"rating"`
    CreatedAt time.Time  `db:"created_at"`
}

type Chat struct {
    ID        int       `db:"id"`
    Name      *string   `db:"name"`
    CreatedAt time.Time `db:"created_at"`
}

type ChatMember struct {
    ChatID int `db:"chat_id"`
    UserID int `db:"user_id"`
}

type Message struct {
    ID        int64      `db:"id"`
    ChatID    int        `db:"chat_id"`
    UserID    *int       `db:"user_id"`
    Text      string     `db:"text"`
    CreatedAt time.Time  `db:"created_at"`
}

type File struct {
    ID        int64   `db:"id"`
    Name      string  `db:"name"`
    FileURL   string  `db:"file_url"`
    MessageID *int64  `db:"message_id"`
    ReviewID  *int64  `db:"review_id"`
}