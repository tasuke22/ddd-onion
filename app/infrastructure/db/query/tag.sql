-- name: UpsertTag :exec
INSERT INTO tags (id,
                  name,
                  created_at,
                  updated_at)
VALUES (sqlc.arg(id),
        sqlc.arg(name),
        NOW(),
        NOW()) ON DUPLICATE KEY
UPDATE
    name = sqlc.arg(name),
    updated_at = NOW();

-- name: FindByNames :many
SELECT id, name
FROM tags
WHERE name IN (sqlc.slice('names'));

