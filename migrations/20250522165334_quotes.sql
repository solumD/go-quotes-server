-- +goose Up
CREATE TABLE quotes (
    id SERIAL PRIMARY KEY,
    quote_text TEXT NOT NULL,
    quote_author VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE quotes;

