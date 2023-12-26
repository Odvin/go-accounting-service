package service

import (
	"context"
	"errors"

	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/google/uuid"
)

var (
	ErrClientCreateProfile      = errors.New("failed to create client profile")
	ErrClientProfileDuplication = errors.New("duplication of the client profile")
)

type Client struct {
	repository db.Store
}

func ClientService(repository db.Store) *Client {
	return &Client{repository: repository}
}

type ClientInfo struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type ClientSub struct {
	ID           uuid.UUID
	HashPassword string
}

func (client *Client) CreateProfile(info ClientInfo, sub ClientSub) (string, error) {
	arg := db.CreateClientProfileParams{
		ID:       sub.ID,
		Adm:      db.AdministrativeStatusAdmActive,
		Kyc:      db.KycStatusKycConfirmed,
		Name:     info.Name,
		Surname:  info.Surname,
		Password: sub.HashPassword,
		Email:    info.Email,
	}

	id, err := client.repository.CreateClientProfile(context.Background(), arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return "", ErrClientProfileDuplication
		}
		return "", ErrClientCreateProfile
	}

	return id.String(), nil
}
