-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS indicators (
                                          username varchar(25),
                                          indicator varchar(255),
                                          score integer,
                                          total integer,
                                          CONSTRAINT check_score_not_exceed_total CHECK ((score <= total))
);;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS indicators;
-- +goose StatementEnd
