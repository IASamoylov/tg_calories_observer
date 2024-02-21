-- +goose Up
-- +goose StatementBegin
create table "user"
(
    telegram_id bigint     not null primary key,
    user_name   text       not null,
    first_name  text       not null,
    last_name   text       not null,
    language    varchar(2) not null,
    agreement   boolean default false
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "user"
-- +goose StatementEnd
