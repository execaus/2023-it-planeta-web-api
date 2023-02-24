-- name: CreateAccount :one
INSERT INTO "Account" (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: IsExistAccountByEmail :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE email=$1
  AND deleted=false
);

-- name: IsExistAccountById :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE id=$1
  AND deleted=false
);

-- name: GetAccount :one
SELECT *
FROM "Account"
WHERE id=$1
AND deleted=false;

-- name: GetAccounts :many
SELECT *
FROM "Account"
WHERE (first_name IS NULL OR lower(first_name) LIKE lower('%' || $1 || '%'))
AND (last_name IS NULL OR lower(last_name) LIKE lower('%' || $2 || '%'))
AND (email IS NULL OR lower(email) LIKE lower('%' || $3 || '%'))
AND deleted=false
ORDER BY id DESC
LIMIT $4 OFFSET $5;

-- name: UpdateAccount :one
UPDATE "Account"
SET first_name=$1, last_name=$2, email=$3, password=$4
WHERE id=$5
AND deleted=false
RETURNING *;

-- name: RemoveAccount :one
UPDATE "Account"
SET deleted=true
WHERE id=$1
RETURNING *;

-- name: GetAnimal :one
SELECT *
FROM "Animal"
WHERE id=$1
AND deleted=false;

-- name: GetAnimalTypesFromAnimal :many
SELECT *
FROM "AnimalToType"
WHERE animal=$1
AND deleted=false;

-- name: GetVisitedLocationFromAnimal :many
SELECT *
FROM "AnimalVisitedLocation"
WHERE animal=$1
AND deleted=false;

-- name: GetLocation :one
SELECT *
FROM "LocationPoint"
WHERE id=$1
AND deleted=false;

-- name: IsExistLocation :one
SELECT EXISTS (
  SELECT 1
  FROM "LocationPoint"
  WHERE id=$1
  AND deleted=false
);
