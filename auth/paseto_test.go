package auth

import (
	"testing"
	"time"

	"github.com/Odvin/go-accounting-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreatePasetoAuditor(t *testing.T) {
	auditor, err := CreatePasetoAuditor(util.RandomString(32))
	require.NoError(t, err)

	sub := uuid.NewString()
	role := util.DepositorRole
	duration := time.Minute

	issued := time.Now()
	expired := issued.Add(duration)

	token, payload, err := auditor.CreateToken(sub, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = auditor.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, sub, payload.Sub)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issued, payload.Issued, time.Second)
	require.WithinDuration(t, expired, payload.Expired, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	auditor, err := CreatePasetoAuditor(util.RandomString(32))
	require.NoError(t, err)

	token, payload, err := auditor.CreateToken(uuid.NewString(), util.DepositorRole, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = auditor.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
