-- name: UpsertSkill :exec
INSERT INTO skills (id,
                    user_id,
                    tag_id,
                    evaluation,
                    years,
                    created_at,
                    updated_at)
VALUES (sqlc.arg(id),
        sqlc.arg(user_id),
        sqlc.arg(tag_id),
        sqlc.arg(evaluation),
        sqlc.arg(years),
        NOW(),
        NOW()) ON DUPLICATE KEY
UPDATE
    user_id = sqlc.arg(user_id),
    tag_id = sqlc.arg(tag_id),
    evaluation = sqlc.arg(evaluation),
    years = sqlc.arg(years),
    updated_at = NOW();

-- name: FindSkillsByUserID :many
SELECT *
FROM skills
WHERE user_id = ?;