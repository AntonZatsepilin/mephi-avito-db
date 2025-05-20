CREATE TABLE user_types (
    id      SERIAL PRIMARY KEY,
    type TEXT NOT NULL UNIQUE CHECK (type IN ('admin','user','moderator'))
);

CREATE TABLE countries (
    id      SERIAL PRIMARY KEY,
    name    TEXT NOT NULL UNIQUE
);

CREATE TABLE cities (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    country_id  INTEGER NOT NULL
        REFERENCES countries(id)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    UNIQUE (country_id, name)
);

CREATE TABLE categories (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    parent_id   INTEGER
        REFERENCES categories(id)
        ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT categories_parent_not_self CHECK (parent_id IS NULL OR parent_id <> id)
);


CREATE TABLE users (
    id               SERIAL PRIMARY KEY,
    username         TEXT NOT NULL UNIQUE,
    email            TEXT NOT NULL UNIQUE,
    phone_number     TEXT,
    password_hash    TEXT NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    profile_image    TEXT,
    rating           NUMERIC(3,2)  CHECK (rating BETWEEN 0 AND 5),
    user_type_id     INTEGER NOT NULL
        REFERENCES user_types(id)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    city_id          INTEGER
        REFERENCES cities(id)
        ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE listings (
    id            BIGSERIAL PRIMARY KEY,
    user_id       INTEGER NOT NULL
        REFERENCES users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    category_id   INTEGER NOT NULL
        REFERENCES categories(id)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    city_id       INTEGER NOT NULL
        REFERENCES cities(id)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    title         TEXT NOT NULL,
    description   TEXT,
    price         NUMERIC(12,2) NOT NULL CHECK (price >= 0),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_active     BOOLEAN NOT NULL DEFAULT TRUE,
    view_count    INTEGER NOT NULL DEFAULT 0 CHECK (view_count >= 0)
);

CREATE TABLE reviews (
    id           BIGSERIAL PRIMARY KEY,
    user_id      INTEGER NOT NULL
        REFERENCES users(id)
        ON UPDATE CASCADE ON DELETE SET NULL,
    listing_id   INTEGER NOT NULL
        REFERENCES listings(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    comment      TEXT,
    rating       SMALLINT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE TABLE chats (
    id           SERIAL PRIMARY KEY,
    name         TEXT,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE chat_members (
    chat_id      INTEGER NOT NULL
        REFERENCES chats(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    user_id      INTEGER NOT NULL
        REFERENCES users(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    PRIMARY KEY (chat_id, user_id)
);

CREATE TABLE messages (
    id           BIGSERIAL PRIMARY KEY,
    chat_id      INTEGER NOT NULL
        REFERENCES chats(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    user_id      INTEGER NOT NULL
        REFERENCES users(id)
        ON UPDATE CASCADE ON DELETE SET NULL,
    text         TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE files (
    id           BIGSERIAL PRIMARY KEY,
    name         TEXT NOT NULL,
    file_url     TEXT NOT NULL,
    message_id   BIGINT NULL
        REFERENCES messages(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    review_id    BIGINT NULL
        REFERENCES reviews(id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT one_attach_only 
      CHECK (
         (message_id IS NOT NULL AND review_id IS NULL)
      OR (message_id IS NULL AND review_id IS NOT NULL)
      )
);

INSERT INTO user_types (type) VALUES ('admin'), ('user'), ('moderator');