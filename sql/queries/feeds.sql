-- name: CreateFeed :one
INSERT INTO feeds (
  id,
  created_at,
  updated_at,
  user_id,
  url,
  name
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds ORDER BY created_at DESC;

-- name: GetFeedsByUserID :many
SELECT * FROM feeds WHERE user_id = $1 ORDER BY created_at DESC;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;