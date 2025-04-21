-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tags (
    id              serial PRIMARY KEY,
    created_at      timestamp NOT NULL DEFAULT now(),
    updated_at      timestamp NOT NULL DEFAULT now(),
    count           integer NOT NULL,
    type            integer NOT NULL,
    name            varchar(100)
);
CREATE UNIQUE INDEX IF NOT EXISTS tags_name_unique_idx ON tags (lower(name));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tags;
-- +goose StatementEnd
