// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: client-profile.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createClientProfile = `-- name: CreateClientProfile :one
INSERT INTO client.profile (
  id,
  adm,
  kyc,
  name,
  surname,
  password,
  email
) VALUES (
  $1, $2, $3, $4, $5, $6, $7 
) RETURNING id
`

type CreateClientProfileParams struct {
	ID       uuid.UUID            `json:"id"`
	Adm      AdministrativeStatus `json:"adm"`
	Kyc      KycStatus            `json:"kyc"`
	Name     string               `json:"name"`
	Surname  string               `json:"surname"`
	Password string               `json:"password"`
	Email    string               `json:"email"`
}

func (q *Queries) CreateClientProfile(ctx context.Context, arg CreateClientProfileParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createClientProfile,
		arg.ID,
		arg.Adm,
		arg.Kyc,
		arg.Name,
		arg.Surname,
		arg.Password,
		arg.Email,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteClientProfile = `-- name: DeleteClientProfile :exec
DELETE FROM client.profile
WHERE id = $1
`

func (q *Queries) DeleteClientProfile(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteClientProfile, id)
	return err
}

const getClientPasswordByEmail = `-- name: GetClientPasswordByEmail :one
SELECT password, password_updated FROM client.profile
WHERE email = $1 LIMIT 1
`

type GetClientPasswordByEmailRow struct {
	Password        string    `json:"password"`
	PasswordUpdated time.Time `json:"password_updated"`
}

func (q *Queries) GetClientPasswordByEmail(ctx context.Context, email string) (GetClientPasswordByEmailRow, error) {
	row := q.db.QueryRow(ctx, getClientPasswordByEmail, email)
	var i GetClientPasswordByEmailRow
	err := row.Scan(&i.Password, &i.PasswordUpdated)
	return i, err
}

const getClientPasswordById = `-- name: GetClientPasswordById :one
SELECT password, password_updated FROM client.profile
WHERE id = $1 LIMIT 1
`

type GetClientPasswordByIdRow struct {
	Password        string    `json:"password"`
	PasswordUpdated time.Time `json:"password_updated"`
}

func (q *Queries) GetClientPasswordById(ctx context.Context, id uuid.UUID) (GetClientPasswordByIdRow, error) {
	row := q.db.QueryRow(ctx, getClientPasswordById, id)
	var i GetClientPasswordByIdRow
	err := row.Scan(&i.Password, &i.PasswordUpdated)
	return i, err
}

const getClientProfile = `-- name: GetClientProfile :one
SELECT id, adm, kyc, name, surname, updated, created, password, email, password_updated FROM client.profile
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetClientProfile(ctx context.Context, id uuid.UUID) (ClientProfile, error) {
	row := q.db.QueryRow(ctx, getClientProfile, id)
	var i ClientProfile
	err := row.Scan(
		&i.ID,
		&i.Adm,
		&i.Kyc,
		&i.Name,
		&i.Surname,
		&i.Updated,
		&i.Created,
		&i.Password,
		&i.Email,
		&i.PasswordUpdated,
	)
	return i, err
}

const updateClientPassword = `-- name: UpdateClientPassword :exec
UPDATE client.profile
SET
  password = $2,
  password_updated = now()
WHERE
  id = $1
RETURNING id, adm, kyc, name, surname, updated, created, password, email, password_updated
`

type UpdateClientPasswordParams struct {
	ID       uuid.UUID `json:"id"`
	Password string    `json:"password"`
}

func (q *Queries) UpdateClientPassword(ctx context.Context, arg UpdateClientPasswordParams) error {
	_, err := q.db.Exec(ctx, updateClientPassword, arg.ID, arg.Password)
	return err
}

const updateClientProfile = `-- name: UpdateClientProfile :exec
UPDATE client.profile
SET
  name = COALESCE($2, name),
  surname = COALESCE($3, surname),
  email = COALESCE($4, email),
  updated = now()
WHERE
  id = $1
`

type UpdateClientProfileParams struct {
	ID      uuid.UUID   `json:"id"`
	Name    pgtype.Text `json:"name"`
	Surname pgtype.Text `json:"surname"`
	Email   pgtype.Text `json:"email"`
}

func (q *Queries) UpdateClientProfile(ctx context.Context, arg UpdateClientProfileParams) error {
	_, err := q.db.Exec(ctx, updateClientProfile,
		arg.ID,
		arg.Name,
		arg.Surname,
		arg.Email,
	)
	return err
}

const updateClientStatus = `-- name: UpdateClientStatus :exec
UPDATE client.profile
SET
  adm = COALESCE($2, adm),
  kyc = COALESCE($3, kyc),
  updated = now()
WHERE
  id = $1
`

type UpdateClientStatusParams struct {
	ID  uuid.UUID                `json:"id"`
	Adm NullAdministrativeStatus `json:"adm"`
	Kyc NullKycStatus            `json:"kyc"`
}

func (q *Queries) UpdateClientStatus(ctx context.Context, arg UpdateClientStatusParams) error {
	_, err := q.db.Exec(ctx, updateClientStatus, arg.ID, arg.Adm, arg.Kyc)
	return err
}
