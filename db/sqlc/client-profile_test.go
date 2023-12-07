package db

import (
	"context"
	"testing"
	"time"

	"github.com/Odvin/go-accounting-service/util"
	"github.com/stretchr/testify/require"
)

func createClientProfileParams() CreateClientProfileParams {
	return CreateClientProfileParams{
		ID:      util.RandomUUID(),
		Adm:     AdministrativeStatus(util.RandomAdmStatus()),
		Kyc:     KycStatus(util.RandomKycStatus()),
		Name:    util.RandomName(),
		Surname: util.RandomSurname(),
	}
}

func TestCreateClientProfile(t *testing.T) {

	arg := createClientProfileParams()

	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, id)
}

func TestGetClientProfile(t *testing.T) {

	arg := createClientProfileParams()

	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotZero(t, id)

	clientProfile, err := testStore.GetClientProfile(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, clientProfile)

	require.NotZero(t, clientProfile.Created)
	require.NotZero(t, clientProfile.Updated)

	require.Equal(t, id, clientProfile.ID)

	require.WithinDuration(t, clientProfile.Created, clientProfile.Updated, time.Second)

	require.Equal(t, arg.Adm, clientProfile.Adm)
	require.Equal(t, arg.Kyc, clientProfile.Kyc)
	require.Equal(t, arg.Name, clientProfile.Name)
	require.Equal(t, arg.Surname, clientProfile.Surname)
}
