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
WHERE (lower(first_name) LIKE lower('%' || $1 || '%') OR first_name IS NOT DISTINCT FROM NULL)
AND (lower(last_name) LIKE lower('%' || $2 || '%') OR last_name IS NOT DISTINCT FROM NULL)
AND (lower(email) LIKE lower('%' || $3 || '%') OR email IS NOT DISTINCT FROM NULL)
AND deleted = FALSE
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

-- name: IsExistAnimalByID :one
SELECT EXISTS (
  SELECT 1
  FROM "Animal"
  WHERE id=$1
  AND deleted=false
);

-- name: CreateAnimal :one
INSERT INTO "Animal"
(chipping_location, "weight", "length", "height", gender, life_status, chipping_date, chipper, death_date, deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetAnimal :one
SELECT *
FROM "Animal"
WHERE id=$1
AND deleted=false;

-- name: GetAnimals :many
SELECT *
FROM "Animal"
WHERE
    ("chipping_date" >= COALESCE($1, "chipping_date"))
    AND ("chipping_date" <= COALESCE($2, "chipping_date"))
    AND ($3 IS NOT DISTINCT FROM "chipper" OR "chipper" = $3)
    AND ($4 IS NOT DISTINCT FROM "chipping_location" OR "chipping_location" = $4)
    AND ($5 IS NOT DISTINCT FROM "life_status" OR "life_status" = $5)
    AND ($6 IS NOT DISTINCT FROM "gender" OR "gender" = $6)
    AND "deleted" = FALSE
ORDER BY "id"
LIMIT $7 OFFSET $8;

-- name: UpdateAnimal :one
UPDATE "Animal"
SET weight=$1
AND length=$2
AND height=$3
AND gender=$4
AND life_status=$5
AND chipper=$6
AND chipping_location=$7
WHERE id=$8
RETURNING *;

-- name: RemoveAnimal :one
UPDATE "Animal"
SET deleted=true
WHERE id=$1
RETURNING *;

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

-- name: BindAnimalTypeToAnimal :one
INSERT INTO "AnimalToType" (animal, animal_type)
VALUES ($1, $2)
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

-- name: IsLinkedAnimalType :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalToType"
  WHERE animal=$1
  AND animal_type=$2
);

-- name: LinkAnimalTypeToAnimal :one
INSERT INTO "AnimalToType" (animal, animal_type)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateAnimalTypeToAnimal :one
UPDATE "AnimalToType"
SET animal_type=$1
WHERE animal=$2
AND animal_type=$3
RETURNING *;

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

-- name: GetVisitedLocation :one
SELECT *
FROM "AnimalVisitedLocation"
WHERE id=$1;

-- name: GetVisitedLocations :many
SELECT *
FROM "AnimalVisitedLocation"
WHERE animal=$1;

-- name: IsExistVisitedLocationByID :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalVisitedLocation"
  WHERE id=$1
  AND deleted=false
);

-- name: IsLinkedAnimalToVisitedLocation :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalVisitedLocation"
  WHERE id=$1
  AND animal=$2
  AND deleted=false
);

-- name: UpdateVisitedLocation :one
UPDATE "AnimalVisitedLocation"
SET location=$1
WHERE id=$2
AND deleted=false
RETURNING *;

-- name: RemoveVisitedLocation :one
UPDATE "AnimalVisitedLocation"
SET deleted=true
WHERE id=$1
RETURNING *;

-- name: GetVisitedLocationList :many
SELECT *
FROM "AnimalVisitedLocation"
WHERE "deleted" = false
    AND "animal" = $1
    AND ("date" >= COALESCE($2::TIMESTAMP, '1970-01-01'::TIMESTAMP))
    AND ("date" <= COALESCE($3::TIMESTAMP, NOW()))
ORDER BY "date" ASC
OFFSET $4
LIMIT $5;
