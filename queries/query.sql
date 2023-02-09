-- name: CreateAccount :one
INSERT INTO "Account" (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: IsExistAccountByEmail :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE email = $1
);

-- name: IsExistAccountById :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE id = $1
);

-- name: GetAccount :one
SELECT * FROM "Account" WHERE id=$1;