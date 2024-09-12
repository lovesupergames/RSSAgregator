-- +goose Up
CREATE TABLE posts(
                     id uuid PRIMARY KEY,
                     created_at timestamp not null ,
                     updated_at timestamp not null ,
                     title varchar(100) not null,
                     url varchar(100) not null unique,
                     description varchar(100),
                     published_at timestamp not null ,
                     feed_id uuid not null,
                     constraint fk_feed_id foreign key (feed_id) references feed(feed_id)
                         on delete cascade
);

-- +goose Down
DROP TABLE posts;