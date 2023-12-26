package service

import (
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/google/uuid"
)

type Manager interface {
	CreateProfile(info ClientInfo) (string, error)
	GetProfile(id uuid.UUID) db.ClientProfile
}
