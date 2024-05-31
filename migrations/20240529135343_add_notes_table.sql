-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE notes (
    Id SERIAL PRIMARY KEY,
    UserId INT NOT NULL,
    Title VARCHAR NOT NULL,
    Text VARCHAR,
    CreatedAt TIMESTAMP NOT NULL,
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE SET NULL ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
TRUNCATE notes;
DROP TABLE notes;
-- +goose StatementEnd

