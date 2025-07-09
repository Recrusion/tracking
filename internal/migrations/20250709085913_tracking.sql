-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_tokens (
    username varchar(25),
    access_token text,
    created timestamp,
    ending timestamp 
);

CREATE TABLE IF NOT EXISTS indicators (
    username varchar(25),
    indicator varchar(255),
    score integer,
    total integer,
    CONSTRAINT check_score_not_exceed_total CHECK ((score <= total))
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    username varchar(25),
    refresh_token text,
    created timestamp,
    ending timestamp
);

CREATE TABLE IF NOT EXISTS users (
    id serial primary key NOT NULL,
    username varchar(25),
    password varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS indicators;
DROP TABLE IF EXISTS access_tokens;
-- +goose StatementEnd