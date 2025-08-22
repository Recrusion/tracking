-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_tokens (
                                             username varchar(25),
                                             access_token text,
                                             created timestamp,
                                             ending timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_tokens;
-- +goose StatementEnd
