package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID      uuid.UUID `json:"id"`
	Sub     string    `json:"sub"`
	Role    string    `json:"role"`
	Issued  time.Time `json:"issued"`
	Expired time.Time `json:"expired"`
}

// CreatePayload creates a new token payload with a specific username and duration
func CreatePayload(sub string, role string, duration time.Duration) (*Payload, error) {
	tokenID := uuid.New()

	payload := &Payload{
		ID:      tokenID,
		Sub:     sub,
		Role:    role,
		Issued:  time.Now(),
		Expired: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expired) {
		return ErrExpiredToken
	}
	return nil
}
