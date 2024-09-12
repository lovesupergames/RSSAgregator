-- name: CreateFeed :one
INSERT INTO feed (id, created_at, updated_at, name, url, userId, feed_id)
VALUES ($1, $2, $3, $4,$5,$6, $7)
RETURNING *;

-- name: GetAllFeed :many
select * from feed;

-- name: CrateFeedFollow :one
select id,feed_id,userId,created_at,updated_at
from feed
where feed_id = $1;

-- name: DeleteFeedFollow :one
DELETE from feed where feed_id = $1
returning *;

-- name: GetAllFeedForUser :many
select * from feed
where userId = $1;

-- name: GetNextFeedsToFetch :many
select * from feed
order by last_fetched_at desc NULLS FIRST
limit $1;

-- name: MarkFeedFetched :one
UPDATE feed
SET created_at =$1, last_fetched_at=$2
where feed_id = $3
returning *;

