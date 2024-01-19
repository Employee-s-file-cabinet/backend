-- +goose Up
-- +goose StatementBegin
BEGIN;
ALTER TABLE users
    ADD COLUMN "mobile_phone_number" varchar NOT NULL,
    ADD COLUMN "office_phone_number" varchar NOT NULL;

UPDATE users SET mobile_phone_number = phone_numbers ->> 'mobile';

ALTER TABLE users
    DROP COLUMN "phone_numbers";
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd