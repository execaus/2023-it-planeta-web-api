-- name: CreateAccount :one
INSERT INTO "Account" (first_name, last_name, email, "password") VALUES ($1, $2, $3, $4) RETURNING *;

-- name: IsExistAccount :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE email = $1
);