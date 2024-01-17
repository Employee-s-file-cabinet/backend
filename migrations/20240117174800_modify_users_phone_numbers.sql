-- +goose Up
-- +goose StatementBegin
BEGIN;
ALTER TABLE users
    ADD COLUMN "mobile_phone_number" varchar NOT NULL,
    ADD COLUMN "office_phone_number" varchar NOT NULL,
    DROP COLUMN "phone_numbers";
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
