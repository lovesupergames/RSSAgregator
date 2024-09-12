-- +goose Up
CREATE TABLE feed(
                      id uuid PRIMARY KEY,
                      created_at timestamp not null ,
                      updated_at timestamp not null ,
                      name varchar(100) not null,
                      url varchar(100) not null unique,
                      userId uuid not null,
                      constraint fk_userId foreign key (userId) references users(id)
                      on delete cascade
);

-- +goose Down
DROP TABLE feed;