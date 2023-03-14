// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package queries

import (
	"context"
	"database/sql"
	"time"
)

const bindAnimalTypeToAnimal = `-- name: BindAnimalTypeToAnimal :one
INSERT INTO "AnimalToType" (animal, animal_type)
VALUES ($1, $2)
RETURNING animal, animal_type
`

type BindAnimalTypeToAnimalParams struct {
	Animal     int64
	AnimalType int64
}

func (q *Queries) BindAnimalTypeToAnimal(ctx context.Context, arg BindAnimalTypeToAnimalParams) (AnimalToType, error) {
	row := q.db.QueryRowContext(ctx, bindAnimalTypeToAnimal, arg.Animal, arg.AnimalType)
	var i AnimalToType
	err := row.Scan(&i.Animal, &i.AnimalType)
	return i, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO "Account" (first_name, last_name, email, password, deleted)
VALUES ($1, $2, $3, $4, false)
RETURNING id, first_name, last_name, email, password, deleted
`

type CreateAccountParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Deleted,
	)
	return i, err
}

const createAnimal = `-- name: CreateAnimal :one
INSERT INTO "Animal"
(chipping_location, "weight", "length", "height", gender, life_status, chipping_date, chipper, death_date, deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, chipping_location, weight, length, height, gender, life_status, chipping_date, chipper, death_date, deleted
`

type CreateAnimalParams struct {
	ChippingLocation int64
	Weight           float64
	Length           float64
	Height           float64
	Gender           string
	LifeStatus       string
	ChippingDate     time.Time
	Chipper          int64
	DeathDate        sql.NullTime
	Deleted          bool
}

func (q *Queries) CreateAnimal(ctx context.Context, arg CreateAnimalParams) (Animal, error) {
	row := q.db.QueryRowContext(ctx, createAnimal,
		arg.ChippingLocation,
		arg.Weight,
		arg.Length,
		arg.Height,
		arg.Gender,
		arg.LifeStatus,
		arg.ChippingDate,
		arg.Chipper,
		arg.DeathDate,
		arg.Deleted,
	)
	var i Animal
	err := row.Scan(
		&i.ID,
		&i.ChippingLocation,
		&i.Weight,
		&i.Length,
		&i.Height,
		&i.Gender,
		&i.LifeStatus,
		&i.ChippingDate,
		&i.Chipper,
		&i.DeathDate,
		&i.Deleted,
	)
	return i, err
}

const createAnimalType = `-- name: CreateAnimalType :one
INSERT INTO "AnimalType" ("value", deleted)
VALUES ($1, false)
RETURNING id, value, deleted
`

func (q *Queries) CreateAnimalType(ctx context.Context, value string) (AnimalType, error) {
	row := q.db.QueryRowContext(ctx, createAnimalType, value)
	var i AnimalType
	err := row.Scan(&i.ID, &i.Value, &i.Deleted)
	return i, err
}

const createLocation = `-- name: CreateLocation :one
INSERT INTO "LocationPoint" (latitude, longitude, deleted)
VALUES ($1, $2, false)
RETURNING id, latitude, longitude, deleted
`

type CreateLocationParams struct {
	Latitude  float64
	Longitude float64
}

func (q *Queries) CreateLocation(ctx context.Context, arg CreateLocationParams) (LocationPoint, error) {
	row := q.db.QueryRowContext(ctx, createLocation, arg.Latitude, arg.Longitude)
	var i LocationPoint
	err := row.Scan(
		&i.ID,
		&i.Latitude,
		&i.Longitude,
		&i.Deleted,
	)
	return i, err
}

const createVisitedLocation = `-- name: CreateVisitedLocation :one
INSERT INTO
"AnimalVisitedLocation" (location, animal, date, deleted)
VALUES ($1, $2, now(), false)
RETURNING id, location, animal, date, deleted
`

type CreateVisitedLocationParams struct {
	Location int64
	Animal   int64
}

func (q *Queries) CreateVisitedLocation(ctx context.Context, arg CreateVisitedLocationParams) (AnimalVisitedLocation, error) {
	row := q.db.QueryRowContext(ctx, createVisitedLocation, arg.Location, arg.Animal)
	var i AnimalVisitedLocation
	err := row.Scan(
		&i.ID,
		&i.Location,
		&i.Animal,
		&i.Date,
		&i.Deleted,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, first_name, last_name, email, password, deleted
FROM "Account"
WHERE id=$1
AND deleted=false
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Deleted,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
SELECT id, first_name, last_name, email, password, deleted
FROM "Account"
WHERE (lower(first_name) LIKE lower('%' || $1 || '%') OR first_name IS NOT DISTINCT FROM NULL)
AND (lower(last_name) LIKE lower('%' || $2 || '%') OR last_name IS NOT DISTINCT FROM NULL)
AND (lower(email) LIKE lower('%' || $3 || '%') OR email IS NOT DISTINCT FROM NULL)
AND deleted = FALSE
ORDER BY id DESC
LIMIT $4 OFFSET $5
`

type GetAccountsParams struct {
	Column1 sql.NullString
	Column2 sql.NullString
	Column3 sql.NullString
	Limit   int32
	Offset  int32
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts,
		arg.Column1,
		arg.Column2,
		arg.Column3,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAnimal = `-- name: GetAnimal :one
SELECT id, chipping_location, weight, length, height, gender, life_status, chipping_date, chipper, death_date, deleted
FROM "Animal"
WHERE id=$1
AND deleted=false
`

func (q *Queries) GetAnimal(ctx context.Context, id int64) (Animal, error) {
	row := q.db.QueryRowContext(ctx, getAnimal, id)
	var i Animal
	err := row.Scan(
		&i.ID,
		&i.ChippingLocation,
		&i.Weight,
		&i.Length,
		&i.Height,
		&i.Gender,
		&i.LifeStatus,
		&i.ChippingDate,
		&i.Chipper,
		&i.DeathDate,
		&i.Deleted,
	)
	return i, err
}

const getAnimalTypeByID = `-- name: GetAnimalTypeByID :one
SELECT id, value, deleted
FROM "AnimalType"
WHERE id=$1
`

func (q *Queries) GetAnimalTypeByID(ctx context.Context, id int64) (AnimalType, error) {
	row := q.db.QueryRowContext(ctx, getAnimalTypeByID, id)
	var i AnimalType
	err := row.Scan(&i.ID, &i.Value, &i.Deleted)
	return i, err
}

const getAnimalTypesByAnimalID = `-- name: GetAnimalTypesByAnimalID :many
SELECT animal, animal_type
FROM "AnimalToType"
WHERE animal=$1
`

func (q *Queries) GetAnimalTypesByAnimalID(ctx context.Context, animal int64) ([]AnimalToType, error) {
	rows, err := q.db.QueryContext(ctx, getAnimalTypesByAnimalID, animal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AnimalToType
	for rows.Next() {
		var i AnimalToType
		if err := rows.Scan(&i.Animal, &i.AnimalType); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAnimals = `-- name: GetAnimals :many
SELECT id, chipping_location, weight, length, height, gender, life_status, chipping_date, chipper, death_date, deleted
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
LIMIT $7 OFFSET $8
`

type GetAnimalsParams struct {
	ChippingDate   time.Time
	ChippingDate_2 time.Time
	Column3        interface{}
	Column4        interface{}
	Column5        interface{}
	Column6        interface{}
	Limit          int32
	Offset         int32
}

func (q *Queries) GetAnimals(ctx context.Context, arg GetAnimalsParams) ([]Animal, error) {
	rows, err := q.db.QueryContext(ctx, getAnimals,
		arg.ChippingDate,
		arg.ChippingDate_2,
		arg.Column3,
		arg.Column4,
		arg.Column5,
		arg.Column6,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Animal
	for rows.Next() {
		var i Animal
		if err := rows.Scan(
			&i.ID,
			&i.ChippingLocation,
			&i.Weight,
			&i.Length,
			&i.Height,
			&i.Gender,
			&i.LifeStatus,
			&i.ChippingDate,
			&i.Chipper,
			&i.DeathDate,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChippingLocation = `-- name: GetChippingLocation :one
SELECT "LocationPoint".id, "LocationPoint".latitude, "LocationPoint".longitude, "LocationPoint".deleted
FROM "Animal"
JOIN "LocationPoint" ON "Animal".chipping_location = "LocationPoint".id AND "Animal".id=$1
`

func (q *Queries) GetChippingLocation(ctx context.Context, id int64) (LocationPoint, error) {
	row := q.db.QueryRowContext(ctx, getChippingLocation, id)
	var i LocationPoint
	err := row.Scan(
		&i.ID,
		&i.Latitude,
		&i.Longitude,
		&i.Deleted,
	)
	return i, err
}

const getCurrentLocation = `-- name: GetCurrentLocation :one
SELECT id, location, animal, date, deleted
FROM "AnimalVisitedLocation"
WHERE animal=$1 ORDER BY "date" DESC
LIMIT 1
`

func (q *Queries) GetCurrentLocation(ctx context.Context, animal int64) (AnimalVisitedLocation, error) {
	row := q.db.QueryRowContext(ctx, getCurrentLocation, animal)
	var i AnimalVisitedLocation
	err := row.Scan(
		&i.ID,
		&i.Location,
		&i.Animal,
		&i.Date,
		&i.Deleted,
	)
	return i, err
}

const getLocation = `-- name: GetLocation :one
SELECT id, latitude, longitude, deleted
FROM "LocationPoint"
WHERE id=$1
AND deleted=false
`

func (q *Queries) GetLocation(ctx context.Context, id int64) (LocationPoint, error) {
	row := q.db.QueryRowContext(ctx, getLocation, id)
	var i LocationPoint
	err := row.Scan(
		&i.ID,
		&i.Latitude,
		&i.Longitude,
		&i.Deleted,
	)
	return i, err
}

const getVisitedLocation = `-- name: GetVisitedLocation :one
SELECT id, location, animal, date, deleted
FROM "AnimalVisitedLocation"
WHERE id=$1
`

func (q *Queries) GetVisitedLocation(ctx context.Context, id int64) (AnimalVisitedLocation, error) {
	row := q.db.QueryRowContext(ctx, getVisitedLocation, id)
	var i AnimalVisitedLocation
	err := row.Scan(
		&i.ID,
		&i.Location,
		&i.Animal,
		&i.Date,
		&i.Deleted,
	)
	return i, err
}

const getVisitedLocationByAnimalID = `-- name: GetVisitedLocationByAnimalID :many
SELECT id, location, animal, date, deleted
FROM "AnimalVisitedLocation"
WHERE animal=$1
AND deleted=false
`

func (q *Queries) GetVisitedLocationByAnimalID(ctx context.Context, animal int64) ([]AnimalVisitedLocation, error) {
	rows, err := q.db.QueryContext(ctx, getVisitedLocationByAnimalID, animal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AnimalVisitedLocation
	for rows.Next() {
		var i AnimalVisitedLocation
		if err := rows.Scan(
			&i.ID,
			&i.Location,
			&i.Animal,
			&i.Date,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVisitedLocationList = `-- name: GetVisitedLocationList :many
SELECT id, location, animal, date, deleted
FROM "AnimalVisitedLocation"
WHERE "deleted" = false
    AND "animal" = $1
    AND ("date" >= COALESCE($2::TIMESTAMP, '1970-01-01'::TIMESTAMP))
    AND ("date" <= COALESCE($3::TIMESTAMP, NOW()))
ORDER BY "date" ASC
OFFSET $4
LIMIT $5
`

type GetVisitedLocationListParams struct {
	Animal  int64
	Column2 time.Time
	Column3 time.Time
	Offset  int32
	Limit   int32
}

func (q *Queries) GetVisitedLocationList(ctx context.Context, arg GetVisitedLocationListParams) ([]AnimalVisitedLocation, error) {
	rows, err := q.db.QueryContext(ctx, getVisitedLocationList,
		arg.Animal,
		arg.Column2,
		arg.Column3,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AnimalVisitedLocation
	for rows.Next() {
		var i AnimalVisitedLocation
		if err := rows.Scan(
			&i.ID,
			&i.Location,
			&i.Animal,
			&i.Date,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVisitedLocations = `-- name: GetVisitedLocations :many
SELECT id, location, animal, date, deleted
FROM "AnimalVisitedLocation"
WHERE animal=$1
`

func (q *Queries) GetVisitedLocations(ctx context.Context, animal int64) ([]AnimalVisitedLocation, error) {
	rows, err := q.db.QueryContext(ctx, getVisitedLocations, animal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AnimalVisitedLocation
	for rows.Next() {
		var i AnimalVisitedLocation
		if err := rows.Scan(
			&i.ID,
			&i.Location,
			&i.Animal,
			&i.Date,
			&i.Deleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isAnimalTypeLinkedAnimal = `-- name: IsAnimalTypeLinkedAnimal :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalToType"
  WHERE animal_type=$1
)
`

func (q *Queries) IsAnimalTypeLinkedAnimal(ctx context.Context, animalType int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isAnimalTypeLinkedAnimal, animalType)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistAccountByEmail = `-- name: IsExistAccountByEmail :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE email=$1
  AND deleted=false
)
`

func (q *Queries) IsExistAccountByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistAccountByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistAccountByID = `-- name: IsExistAccountByID :one
SELECT EXISTS (
  SELECT 1
  FROM "Account"
  WHERE id=$1
  AND deleted=false
)
`

func (q *Queries) IsExistAccountByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistAccountByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistAnimalByID = `-- name: IsExistAnimalByID :one
SELECT EXISTS (
  SELECT 1
  FROM "Animal"
  WHERE id=$1
  AND deleted=false
)
`

func (q *Queries) IsExistAnimalByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistAnimalByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistAnimalTypeByID = `-- name: IsExistAnimalTypeByID :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalType"
  WHERE id=$1
  AND deleted=false
)
`

func (q *Queries) IsExistAnimalTypeByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistAnimalTypeByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistAnimalTypeByType = `-- name: IsExistAnimalTypeByType :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalType"
  WHERE "value"=$1
  AND deleted=false
)
`

func (q *Queries) IsExistAnimalTypeByType(ctx context.Context, value string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistAnimalTypeByType, value)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistLocationByCoordinates = `-- name: IsExistLocationByCoordinates :one
SELECT EXISTS (
  SELECT 1
  FROM "LocationPoint"
  WHERE latitude=$1
  AND longitude=$2
  AND deleted=false
)
`

type IsExistLocationByCoordinatesParams struct {
	Latitude  float64
	Longitude float64
}

func (q *Queries) IsExistLocationByCoordinates(ctx context.Context, arg IsExistLocationByCoordinatesParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistLocationByCoordinates, arg.Latitude, arg.Longitude)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistLocationByID = `-- name: IsExistLocationByID :one
SELECT EXISTS (
  SELECT 1
  FROM "LocationPoint"
  WHERE id=$1
  AND deleted=false
)
`

func (q *Queries) IsExistLocationByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistLocationByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isExistVisitedLocationByID = `-- name: IsExistVisitedLocationByID :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalVisitedLocation"
  WHERE id=$1
  AND deleted=false
)
`

func (q *Queries) IsExistVisitedLocationByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isExistVisitedLocationByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isLinkedAnimalToVisitedLocation = `-- name: IsLinkedAnimalToVisitedLocation :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalVisitedLocation"
  WHERE id=$1
  AND animal=$2
  AND deleted=false
)
`

type IsLinkedAnimalToVisitedLocationParams struct {
	ID     int64
	Animal int64
}

func (q *Queries) IsLinkedAnimalToVisitedLocation(ctx context.Context, arg IsLinkedAnimalToVisitedLocationParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isLinkedAnimalToVisitedLocation, arg.ID, arg.Animal)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isLinkedAnimalType = `-- name: IsLinkedAnimalType :one
SELECT EXISTS (
  SELECT 1
  FROM "AnimalToType"
  WHERE animal=$1
  AND animal_type=$2
)
`

type IsLinkedAnimalTypeParams struct {
	Animal     int64
	AnimalType int64
}

func (q *Queries) IsLinkedAnimalType(ctx context.Context, arg IsLinkedAnimalTypeParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isLinkedAnimalType, arg.Animal, arg.AnimalType)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isLocationChippingAnimal = `-- name: IsLocationChippingAnimal :one
SELECT EXISTS(SELECT 1 FROM "Animal" WHERE "Animal".chipping_location = "LocationPoint".id)
FROM "LocationPoint"
WHERE "LocationPoint".id=$1
`

func (q *Queries) IsLocationChippingAnimal(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isLocationChippingAnimal, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isLocationVisitedAnimal = `-- name: IsLocationVisitedAnimal :one
SELECT EXISTS(SELECT 1 FROM "AnimalVisitedLocation" WHERE "AnimalVisitedLocation".location = "LocationPoint".id)
FROM "LocationPoint"
WHERE "LocationPoint".id=$1
`

func (q *Queries) IsLocationVisitedAnimal(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, isLocationVisitedAnimal, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const linkAnimalTypeToAnimal = `-- name: LinkAnimalTypeToAnimal :one
INSERT INTO "AnimalToType" (animal, animal_type)
VALUES ($1, $2)
RETURNING animal, animal_type
`

type LinkAnimalTypeToAnimalParams struct {
	Animal     int64
	AnimalType int64
}

func (q *Queries) LinkAnimalTypeToAnimal(ctx context.Context, arg LinkAnimalTypeToAnimalParams) (AnimalToType, error) {
	row := q.db.QueryRowContext(ctx, linkAnimalTypeToAnimal, arg.Animal, arg.AnimalType)
	var i AnimalToType
	err := row.Scan(&i.Animal, &i.AnimalType)
	return i, err
}

const removeAccount = `-- name: RemoveAccount :one
UPDATE "Account"
SET deleted=true
WHERE id=$1
RETURNING id, first_name, last_name, email, password, deleted
`

func (q *Queries) RemoveAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, removeAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Deleted,
	)
	return i, err
}

const removeAnimal = `-- name: RemoveAnimal :one
UPDATE "Animal"
SET deleted=true
WHERE id=$1
RETURNING id, chipping_location, weight, length, height, gender, life_status, chipping_date, chipper, death_date, deleted
`

func (q *Queries) RemoveAnimal(ctx context.Context, id int64) (Animal, error) {
	row := q.db.QueryRowContext(ctx, removeAnimal, id)
	var i Animal
	err := row.Scan(
		&i.ID,
		&i.ChippingLocation,
		&i.Weight,
		&i.Length,
		&i.Height,
		&i.Gender,
		&i.LifeStatus,
		&i.ChippingDate,
		&i.Chipper,
		&i.DeathDate,
		&i.Deleted,
	)
	return i, err
}

const removeAnimalType = `-- name: RemoveAnimalType :one
UPDATE "AnimalType"
SET deleted=true
WHERE id=$1
RETURNING id, value, deleted
`

func (q *Queries) RemoveAnimalType(ctx context.Context, id int64) (AnimalType, error) {
	row := q.db.QueryRowContext(ctx, removeAnimalType, id)
	var i AnimalType
	err := row.Scan(&i.ID, &i.Value, &i.Deleted)
	return i, err
}

const removeAnimalTypeToAnimal = `-- name: RemoveAnimalTypeToAnimal :one
DELETE FROM "AnimalToType"
WHERE animal=$1
AND animal_type=$2
RETURNING animal, animal_type
`

type RemoveAnimalTypeToAnimalParams struct {
	Animal     int64
	AnimalType int64
}

func (q *Queries) RemoveAnimalTypeToAnimal(ctx context.Context, arg RemoveAnimalTypeToAnimalParams) (AnimalToType, error) {
	row := q.db.QueryRowContext(ctx, removeAnimalTypeToAnimal, arg.Animal, arg.AnimalType)
	var i AnimalToType
	err := row.Scan(&i.Animal, &i.AnimalType)
	return i, err
}

const removeLocation = `-- name: RemoveLocation :one
UPDATE "LocationPoint"
SET deleted=true
WHERE id=$1
RETURNING id, latitude, longitude, deleted
`

func (q *Queries) RemoveLocation(ctx context.Context, id int64) (LocationPoint, error) {
	row := q.db.QueryRowContext(ctx, removeLocation, id)
	var i LocationPoint
	err := row.Scan(
		&i.ID,
		&i.Latitude,
		&i.Longitude,
		&i.Deleted,
	)
	return i, err
}

const removeVisitedLocation = `-- name: RemoveVisitedLocation :one
UPDATE "AnimalVisitedLocation"
SET deleted=true
WHERE id=$1
RETURNING id, location, animal, date, deleted
`

func (q *Queries) RemoveVisitedLocation(ctx context.Context, id int64) (AnimalVisitedLocation, error) {
	row := q.db.QueryRowContext(ctx, removeVisitedLocation, id)
	var i AnimalVisitedLocation
	err := row.Scan(
		&i.ID,
		&i.Location,
		&i.Animal,
		&i.Date,
		&i.Deleted,
	)
	return i, err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE "Account"
SET first_name=$1, last_name=$2, email=$3, password=$4
WHERE id=$5
AND deleted=false
RETURNING id, first_name, last_name, email, password, deleted
`

type UpdateAccountParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	ID        int64
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.ID,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Deleted,
	)
	return i, err
}

const updateAnimal = `-- name: UpdateAnimal :one
UPDATE "Animal"
SET weight=$1
AND length=$2
AND height=$3
AND gender=$4
AND life_status=$5
AND chipper=$6
AND chipping_location=$7
AND death_date=$8
WHERE id=$9
RETURNING id, chipping_location, weight, length, height, gender, life_status, chipping_date, chipper, death_date, deleted
`

type UpdateAnimalParams struct {
	Weight           float64
	Length           float64
	Height           float64
	Gender           string
	LifeStatus       string
	Chipper          int64
	ChippingLocation int64
	DeathDate        sql.NullTime
	ID               int64
}

func (q *Queries) UpdateAnimal(ctx context.Context, arg UpdateAnimalParams) (Animal, error) {
	row := q.db.QueryRowContext(ctx, updateAnimal,
		arg.Weight,
		arg.Length,
		arg.Height,
		arg.Gender,
		arg.LifeStatus,
		arg.Chipper,
		arg.ChippingLocation,
		arg.DeathDate,
		arg.ID,
	)
	var i Animal
	err := row.Scan(
		&i.ID,
		&i.ChippingLocation,
		&i.Weight,
		&i.Length,
		&i.Height,
		&i.Gender,
		&i.LifeStatus,
		&i.ChippingDate,
		&i.Chipper,
		&i.DeathDate,
		&i.Deleted,
	)
	return i, err
}

const updateAnimalType = `-- name: UpdateAnimalType :one
UPDATE "AnimalType"
SET "value"=$1
WHERE id=$2
RETURNING id, value, deleted
`

type UpdateAnimalTypeParams struct {
	Value string
	ID    int64
}

func (q *Queries) UpdateAnimalType(ctx context.Context, arg UpdateAnimalTypeParams) (AnimalType, error) {
	row := q.db.QueryRowContext(ctx, updateAnimalType, arg.Value, arg.ID)
	var i AnimalType
	err := row.Scan(&i.ID, &i.Value, &i.Deleted)
	return i, err
}

const updateAnimalTypeToAnimal = `-- name: UpdateAnimalTypeToAnimal :one
UPDATE "AnimalToType"
SET animal_type=$1
WHERE animal=$2
AND animal_type=$3
RETURNING animal, animal_type
`

type UpdateAnimalTypeToAnimalParams struct {
	AnimalType   int64
	Animal       int64
	AnimalType_2 int64
}

func (q *Queries) UpdateAnimalTypeToAnimal(ctx context.Context, arg UpdateAnimalTypeToAnimalParams) (AnimalToType, error) {
	row := q.db.QueryRowContext(ctx, updateAnimalTypeToAnimal, arg.AnimalType, arg.Animal, arg.AnimalType_2)
	var i AnimalToType
	err := row.Scan(&i.Animal, &i.AnimalType)
	return i, err
}

const updateLocation = `-- name: UpdateLocation :one
UPDATE "LocationPoint"
SET latitude=$1, longitude=$2
WHERE id=$3
RETURNING id, latitude, longitude, deleted
`

type UpdateLocationParams struct {
	Latitude  float64
	Longitude float64
	ID        int64
}

func (q *Queries) UpdateLocation(ctx context.Context, arg UpdateLocationParams) (LocationPoint, error) {
	row := q.db.QueryRowContext(ctx, updateLocation, arg.Latitude, arg.Longitude, arg.ID)
	var i LocationPoint
	err := row.Scan(
		&i.ID,
		&i.Latitude,
		&i.Longitude,
		&i.Deleted,
	)
	return i, err
}

const updateVisitedLocation = `-- name: UpdateVisitedLocation :one
UPDATE "AnimalVisitedLocation"
SET location=$1
WHERE id=$2
AND deleted=false
RETURNING id, location, animal, date, deleted
`

type UpdateVisitedLocationParams struct {
	Location int64
	ID       int64
}

func (q *Queries) UpdateVisitedLocation(ctx context.Context, arg UpdateVisitedLocationParams) (AnimalVisitedLocation, error) {
	row := q.db.QueryRowContext(ctx, updateVisitedLocation, arg.Location, arg.ID)
	var i AnimalVisitedLocation
	err := row.Scan(
		&i.ID,
		&i.Location,
		&i.Animal,
		&i.Date,
		&i.Deleted,
	)
	return i, err
}
