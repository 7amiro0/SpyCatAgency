-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table if not exists cat
(
    id serial primary key,
    name varchar(30) not null,
    salary int not null,
    experience int not null,
    bread varchar(30) not null
);

create table if not exists mission
(
    id serial primary key,
    catID int,
    complete boolean not null,
    targets varchar(20)[] not null
);

create table if not exists targets
(
    name varchar(20) unique,
    country varchar(20),
    notes text,
    complete boolean
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
