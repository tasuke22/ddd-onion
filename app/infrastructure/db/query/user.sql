-- name: UpsertUser :exec
INSERT INTO users (id,
                   email,
                   password,
                   name,
                   profile,
                   created_at,
                   updated_at)
VALUES (sqlc.arg(id),
        sqlc.arg(email),
        sqlc.arg(password),
        sqlc.arg(name),
        sqlc.arg(profile),
        NOW(),
        NOW()) ON DUPLICATE KEY
UPDATE
    email = sqlc.arg(email),
    password = sqlc.arg(password),
    name = sqlc.arg(name),
    profile = sqlc.arg(profile),
    updated_at = NOW();

-- name: FindByName :one
SELECT *
FROM users
WHERE name = ?;
