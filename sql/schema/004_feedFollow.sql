-- +goose Up
ALTER TABLE feed
ADD COLUMN feed_id
uuid not null default gen_random_uuid();


-- +goose Down
ALTER TABLE feed
DROP COLUMN feed_id;