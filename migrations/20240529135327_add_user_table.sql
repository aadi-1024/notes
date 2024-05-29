-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users (
    Id SERIAL PRIMARY KEY,
    Username VARCHAR NOT NULL UNIQUE,
    Email VARCHAR NOT NULL UNIQUE,
    Password VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
