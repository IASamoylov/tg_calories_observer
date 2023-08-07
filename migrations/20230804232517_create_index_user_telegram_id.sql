-- +goose Up
-- +goose StatementBegin
-- +goose NO TRANSACTION
create unique index concurrently idx_user_telegram_id
    on "user" (telegram_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index concurrently idx_user_telegram_id
-- +goose StatementEnd
