-- +goose Up
-- +goose StatementBegin
BEGIN;

CREATE TABLE IF NOT EXISTS "policies"
(
    "id"    bigserial PRIMARY KEY,
    "ptype" varchar NOT NULL,
    "v0"    varchar,
    "v1"    varchar,
    "v2"    varchar,
    "v3"    varchar,
    "v4"    varchar,
    "v5"    varchar
);

COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;

DROP TABLE IF EXISTS policies;

COMMIT;
-- +goose StatementEnd
