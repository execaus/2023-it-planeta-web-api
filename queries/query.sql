-- name: CreateAccount :one
INSERT INTO "Account" (first_name, last_name, email, password, deleted)
VALUES ($1, $2, $3, $4, false)
RETURNING *;

-- name: IsExistAccountByEmail :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE email=$1
  AND deleted=false
);

-- name: IsExistAccountByID :one
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

-- name: IsExistAnimalTypeByID :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalType"
  WHERE id=$1
  AND deleted=false
);

-- name: IsExistAnimalTypeByType :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalType"
  WHERE "value"=$1
  AND deleted=false
);

-- name: CreateAnimalType :one
INSERT INTO "AnimalType" ("value", deleted)
VALUES ($1, false)
RETURNING *;

-- name: UpdateAnimalType :one
UPDATE "AnimalType"
SET "value"=$1
WHERE id=$2
RETURNING *;

-- name: GetAnimalTypeByID :one
SELECT *
FROM "AnimalType"
WHERE id=$1;

-- name: GetAnimalTypesByAnimalID :many
SELECT *
FROM "AnimalToType"
WHERE animal=$1;

-- name: IsAnimalTypeLinkedAnimal :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalToType"
  WHERE animal_type=$1
);

-- name: RemoveAnimalType :one
UPDATE "AnimalType"
SET deleted=true
WHERE id=$1
RETURNING *;

-- name: GetVisitedLocationByAnimalID :many
SELECT *
FROM "AnimalVisitedLocation"
WHERE animal=$1
AND deleted=false;

-- name: GetLocation :one
SELECT *
FROM "LocationPoint"
WHERE id=$1
AND deleted=false;

-- name: IsExistLocationByID :one
SELECT EXISTS (
  SELECT 1
  FROM "LocationPoint"
  WHERE id=$1
  AND deleted=false
);

-- name: IsExistLocationByCoordinates :one
SELECT EXISTS (
  SELECT 1
  FROM "LocationPoint"
  WHERE latitude=$1
  AND longitude=$2
  AND deleted=false
);

-- name: CreateLocation :one
INSERT INTO "LocationPoint" (latitude, longitude, deleted)
VALUES ($1, $2, false)
RETURNING *;

-- name: UpdateLocation :one
UPDATE "LocationPoint"
SET latitude=$1, longitude=$2
WHERE id=$3
RETURNING *;

-- name: IsLocationVisitedAnimal :one
SELECT EXISTS(SELECT 1 FROM "AnimalVisitedLocation" WHERE "AnimalVisitedLocation".location = "LocationPoint".id)
FROM "LocationPoint"
WHERE "LocationPoint".id=$1;

-- name: IsLocationChippingAnimal :one
SELECT EXISTS(SELECT 1 FROM "Animal" WHERE "Animal".chipping_location = "LocationPoint".id)
FROM "LocationPoint"
WHERE "LocationPoint".id=$1;

-- name: RemoveLocation :one
UPDATE "LocationPoint"
SET deleted=true
WHERE id=$1
RETURNING *;

-- name: GetChippingLocation :one
SELECT "LocationPoint".*
FROM "Animal"
JOIN "LocationPoint" ON "Animal".chipping_location = "LocationPoint".id AND "Animal".id=$1;

-- name: GetCurrentLocation :one
SELECT *
FROM "AnimalVisitedLocation"
WHERE animal=$1 ORDER BY "date" DESC
LIMIT 1;

-- name: CreateVisitedLocation :one
INSERT INTO
"AnimalVisitedLocation" (location, animal, date, deleted)
VALUES ($1, $2, now(), false)
RETURNING *;
