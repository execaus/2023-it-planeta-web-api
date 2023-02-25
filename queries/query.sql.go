// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package queries

import (
	"context"
	"database/sql"
)

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
WHERE (first_name IS NULL OR lower(first_name) LIKE lower('%' || $1 || '%'))
AND (last_name IS NULL OR lower(last_name) LIKE lower('%' || $2 || '%'))
AND (email IS NULL OR lower(email) LIKE lower('%' || $3 || '%'))
AND deleted=false
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
