package auth

import (
	"time"

	"github.com/google/uuid"
)

// Authenticator is an interface for managing tokens
type Authenticator interface {
	// CreateToken creates a new token for a specific client sub and duration
	CreateToken(sub uuid.UUID, role string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
