package db

import (
	"context"
	"testing"
	"time"

	"github.com/Odvin/go-accounting-service/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createClientProfileParams(t *testing.T) (CreateClientProfileParams, string) {

	password := util.RandomUUID().String()
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	return CreateClientProfileParams{
		ID:       util.RandomUUID(),
		Adm:      AdministrativeStatus(util.RandomAdmStatus()),
		Kyc:      KycStatus(util.RandomKycStatus()),
		Name:     util.RandomName(),
		Surname:  util.RandomSurname(),
		Password: hashedPassword,
		Email:    util.RandomEmail(),
	}, password
}

func TestCreateClientProfile(t *testing.T) {

	arg, _ := createClientProfileParams(t)

	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, id)
}

func TestGetClientProfile(t *testing.T) {

	arg, password := createClientProfileParams(t)
	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotZero(t, id)

	clientProfile, err := testStore.GetClientProfile(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, clientProfile)

	require.NotZero(t, clientProfile.Created)
	require.NotZero(t, clientProfile.Updated)
	require.NotZero(t, clientProfile.PasswordUpdated)

	require.Equal(t, id, clientProfile.ID)

	require.WithinDuration(t, clientProfile.Created, clientProfile.Updated, time.Second)

	require.Equal(t, arg.Adm, clientProfile.Adm)
	require.Equal(t, arg.Kyc, clientProfile.Kyc)
	require.Equal(t, arg.Name, clientProfile.Name)
	require.Equal(t, arg.Surname, clientProfile.Surname)
	require.Equal(t, arg.Email, clientProfile.Email)

	err = util.CheckPassword(password, clientProfile.Password)
	require.NoError(t, err)
}

func TestDeleteClientProfile(t *testing.T) {
	arg, _ := createClientProfileParams(t)
	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	err = testStore.DeleteClientProfile(context.Background(), id)
	require.NoError(t, err)

	clientProfile, err := testStore.GetClientProfile(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, clientProfile)
}

func TestGetClientProfileByEmail(t *testing.T) {
	arg, password := createClientProfileParams(t)

	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	clientProfile, err := testStore.GetClientProfileByEmail(context.Background(), arg.Email)
	require.NoError(t, err)
	require.NotEmpty(t, clientProfile)

	require.NotZero(t, clientProfile.Created)
	require.NotZero(t, clientProfile.Updated)
	require.NotZero(t, clientProfile.PasswordUpdated)

	require.Equal(t, id, clientProfile.ID)

	require.WithinDuration(t, clientProfile.Created, clientProfile.Updated, time.Second)

	require.Equal(t, arg.Adm, clientProfile.Adm)
	require.Equal(t, arg.Kyc, clientProfile.Kyc)
	require.Equal(t, arg.Name, clientProfile.Name)
	require.Equal(t, arg.Surname, clientProfile.Surname)
	require.Equal(t, arg.Email, clientProfile.Email)

	err = util.CheckPassword(password, clientProfile.Password)
	require.NoError(t, err)
}

func TestGetClientPasswordById(t *testing.T) {
	arg, password := createClientProfileParams(t)

	id, err := testStore.CreateClientProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	clientPassword, err := testStore.GetClientPasswordById(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, clientPassword)

	err = util.CheckPassword(password, clientPassword.Password)
	require.NoError(t, err)
}

func TestUpdateClientProfileName(t *testing.T) {
	arg, _ := createClientProfileParams(t)

	id, err := testStore.CreateClientProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	newArg, _ := createClientProfileParams(t)

	err = testStore.UpdateClientProfile(context.Background(), UpdateClientProfileParams{ID: id, Name: pgtype.Text{
		String: newArg.Name,
		Valid:  true,
	}})
	require.NoError(t, err)

	clientProfile, err := testStore.GetClientProfile(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, clientProfile)

	require.Equal(t, id, clientProfile.ID)
	require.Equal(t, arg.Adm, clientProfile.Adm)
	require.Equal(t, arg.Kyc, clientProfile.Kyc)
	require.Equal(t, newArg.Name, clientProfile.Name)
	require.Equal(t, arg.Surname, clientProfile.Surname)
	require.Equal(t, arg.Email, clientProfile.Email)
	require.WithinDuration(t, clientProfile.Updated, time.Now(), time.Minute)
}

func TestUpdateClientProfileSurname(t *testing.T) {
	arg, _ := createClientProfileParams(t)

	id, err := testStore.CreateClientProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	newArg, _ := createClientProfileParams(t)

	err = testStore.UpdateClientProfile(context.Background(), UpdateClientProfileParams{ID: id, Surname: pgtype.Text{
		String: newArg.Surname,
		Valid:  true,
	}})
	require.NoError(t, err)

	clientProfile, err := testStore.GetClientProfile(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, clientProfile)

	require.Equal(t, id, clientProfile.ID)
	require.Equal(t, arg.Adm, clientProfile.Adm)
	require.Equal(t, arg.Kyc, clientProfile.Kyc)
	require.Equal(t, arg.Name, clientProfile.Name)
	require.Equal(t, newArg.Surname, clientProfile.Surname)
	require.Equal(t, arg.Email, clientProfile.Email)
	require.WithinDuration(t, clientProfile.Updated, time.Now(), time.Minute)
}

func TestUpdateClientProfileEmail(t *testing.T) {
	arg, _ := createClientProfileParams(t)

	id, err := testStore.CreateClientProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	newArg, _ := createClientProfileParams(t)

	err = testStore.UpdateClientProfile(context.Background(), UpdateClientProfileParams{ID: id, Email: pgtype.Text{
		String: newArg.Email,
		Valid:  true,
	}})
	require.NoError(t, err)

	clientProfile, err := testStore.GetClientProfile(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, clientProfile)

	require.Equal(t, id, clientProfile.ID)
	require.Equal(t, arg.Adm, clientProfile.Adm)
	require.Equal(t, arg.Kyc, clientProfile.Kyc)
	require.Equal(t, arg.Name, clientProfile.Name)
	require.Equal(t, arg.Surname, clientProfile.Surname)
	require.Equal(t, newArg.Email, clientProfile.Email)
	require.WithinDuration(t, clientProfile.Updated, time.Now(), time.Minute)
}

func TestUpdateClientProfilePassword(t *testing.T) {
	arg, _ := createClientProfileParams(t)

	id, err := testStore.CreateClientProfile(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	newArg, newPassword := createClientProfileParams(t)

	err = testStore.UpdateClientPassword(context.Background(), UpdateClientPasswordParams{ID: id, Password: newArg.Password})
	require.NoError(t, err)

	clientProfile, err := testStore.GetClientProfile(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, clientProfile)
	require.Equal(t, id, clientProfile.ID)

	err = util.CheckPassword(newPassword, clientProfile.Password)

	require.NoError(t, err)
	require.WithinDuration(t, clientProfile.PasswordUpdated, time.Now(), time.Minute)
}
