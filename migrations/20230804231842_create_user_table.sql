-- +goose Up
-- +goose StatementBegin
create table "user"
(
    id          bigserial  not null primary key,
    telegram_id bigint     not null,
    user_name   text       not null,
    first_name  text       not null,
    last_name   text       not null,
    language    varchar(2) not null
);

comment on table "user" is 'the table stored information about telegram users the application will work with PK so as not to be tied to telegram ID';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "user"
-- +goose StatementEnd
