CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name  VARCHAR(100) NOT NULL,
    username   VARCHAR(100) UNIQUE NOT NULL,
    phone      VARCHAR(20),
    email      VARCHAR(150) UNIQUE NOT NULL,
    age        INT CHECK (age >= 0),
    password   TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);