-- +goose Up
CREATE TABLE users (
    Id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    PasswordHash VARCHAR(60) NOT NULL
);

-- +goose Down
DROP TABLE users;