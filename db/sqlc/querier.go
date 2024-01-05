// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateClientProfile(ctx context.Context, arg CreateClientProfileParams) (uuid.UUID, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (ClientSession, error)
	DeleteClientProfile(ctx context.Context, id uuid.UUID) error
	GetClientPasswordById(ctx context.Context, id uuid.UUID) (GetClientPasswordByIdRow, error)
	GetClientProfile(ctx context.Context, id uuid.UUID) (ClientProfile, error)
	GetClientProfileByEmail(ctx context.Context, email string) (ClientProfile, error)
	GetSession(ctx context.Context, id uuid.UUID) (ClientSession, error)
	UpdateClientPassword(ctx context.Context, arg UpdateClientPasswordParams) error
	UpdateClientProfile(ctx context.Context, arg UpdateClientProfileParams) error
	UpdateClientStatus(ctx context.Context, arg UpdateClientStatusParams) error
}

var _ Querier = (*Queries)(nil)
