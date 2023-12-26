package service

import (
	"database/sql"
	"testing"

	mockdb "github.com/Odvin/go-accounting-service/db/mock"
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestClient(t *testing.T) {

	t.Run("profile created successfully", func(t *testing.T) {
		info, sub, arg := randomClientProfile()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := mockdb.NewMockStore(ctrl)
		client := ClientService(repository)

		repository.EXPECT().
			CreateClientProfile(gomock.Any(), gomock.Eq(arg)).
			Times(1).
			Return(arg.ID, nil)

		got, err := client.CreateProfile(info, sub)
		require.NoError(t, err)

		want := arg.ID.String()
		require.Equal(t, got, want)
	})

	t.Run("profile creation failed", func(t *testing.T) {
		info, sub, _ := randomClientProfile()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := mockdb.NewMockStore(ctrl)
		client := ClientService(repository)

		repository.EXPECT().
			CreateClientProfile(gomock.Any(), gomock.Any()).
			Times(1).
			Return(uuid.New(), sql.ErrConnDone)

		_, err := client.CreateProfile(info, sub)
		require.Equal(t, err, ErrClientCreateProfile)
	})

	t.Run("profile duplication", func(t *testing.T) {
		info, sub, _ := randomClientProfile()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := mockdb.NewMockStore(ctrl)
		client := ClientService(repository)

		repository.EXPECT().
			CreateClientProfile(gomock.Any(), gomock.Any()).
			Times(1).
			Return(uuid.New(), db.ErrUniqueViolation)

		_, err := client.CreateProfile(info, sub)
		require.Equal(t, err, ErrClientProfileDuplication)
	})
}

func randomClientProfile() (ClientInfo, ClientSub, db.CreateClientProfileParams) {
	password := uuid.New().String()
	hashedPassword, _ := util.HashPassword(password)
	id := uuid.New()

	info := ClientInfo{
		Name:     util.RandomName(),
		Surname:  util.RandomSurname(),
		Password: password,
		Email:    util.RandomEmail(),
	}

	sub := ClientSub{
		ID:           id,
		HashPassword: hashedPassword,
	}

	arg := db.CreateClientProfileParams{
		ID:       id,
		Adm:      db.AdministrativeStatusAdmActive,
		Kyc:      db.KycStatusKycConfirmed,
		Name:     info.Name,
		Surname:  info.Surname,
		Password: hashedPassword,
		Email:    info.Email,
	}

	return info, sub, arg
}
