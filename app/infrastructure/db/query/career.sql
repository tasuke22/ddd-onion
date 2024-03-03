-- name: UpsertCareer :exec
INSERT INTO careers (id,
                     user_id,
                     detail,
                     start_year,
                     end_year,
                     created_at,
                     updated_at)
VALUES (sqlc.arg(id),
        sqlc.arg(user_id),
        sqlc.arg(detail),
        sqlc.arg(start_year),
        sqlc.arg(end_year),
        NOW(),
        NOW()) ON DUPLICATE KEY
UPDATE
    user_id = sqlc.arg(user_id),
    detail = sqlc.arg(detail),
    start_year = sqlc.arg(start_year),
    end_year = sqlc.arg(end_year),
    updated_at = NOW();

-- name: FindCareersByUserID :many
SELECT *
FROM careers
WHERE user_id = ?;