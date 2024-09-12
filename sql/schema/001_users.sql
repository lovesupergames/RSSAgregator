-- +goose Up
CREATE TABLE users(
    id uuid PRIMARY KEY,
    created_at timestamp not null ,
    updated_at timestamp not null ,
    name varchar(100) not null
);

-- +goose Down
DROP TABLE users;