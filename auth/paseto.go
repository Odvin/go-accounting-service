package auth

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoAuditor struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func CreatePasetoAuditor(symmetricKey string) (Authenticator, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	auditor := &PasetoAuditor{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return auditor, nil
}

func (auditor *PasetoAuditor) CreateToken(sub uuid.UUID, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := CreatePayload(sub, role, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := auditor.paseto.Encrypt(auditor.symmetricKey, payload, nil)
	return token, payload, err
}

func (auditor *PasetoAuditor) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := auditor.paseto.Decrypt(token, auditor.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
