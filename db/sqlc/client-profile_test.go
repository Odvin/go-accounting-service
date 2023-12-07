package db

import (
	"context"
	"testing"

	"github.com/Odvin/go-accounting-service/util"
	"github.com/stretchr/testify/require"
)

func TestCreateClientProfile(t *testing.T) {

	arg := CreateClientProfileParams{
		ID:      util.RandomUUID(),
		Adm:     AdministrativeStatus(util.RandomAdmStatus()),
		Kyc:     KycStatus(util.RandomKycStatus()),
		Name:    util.RandomName(),
		Surname: util.RandomSurname(),
	}

	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, id)
}
