-- name: CreateUser :one
INSERT INTO "users"(
    email, password, username
)VALUES (
            $1,$2,$3
        ) RETURNING *;

-- name: GetUserbyEmail :one
SELECT * FROM "users"
WHERE email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM "users"
WHERE id = $1;