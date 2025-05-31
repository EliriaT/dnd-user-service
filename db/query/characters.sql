-- name: GetCharactersByUserID :many
SELECT * FROM "characters"
WHERE user_id = $1;

-- name: CreateCharacter :one
INSERT INTO "characters" (name, user_id, traits)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: GetCharacterByID :one
SELECT * FROM "characters"
WHERE id = $1;