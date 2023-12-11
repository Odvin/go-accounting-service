package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/Odvin/go-accounting-service/db/mock"
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetClientProfile(t *testing.T) {
	clientProfile := randomClientProfile()
	profileId := clientProfile.ID.String()

	testCases := []struct {
		name          string
		profileID     string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			profileID: profileId,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClientProfile(gomock.Any(), gomock.Eq(clientProfile.ID)).
					Times(1).
					Return(clientProfile, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchClientProfile(t, recorder.Body, clientProfile)
			},
		},
		{
			name:      "InternalError",
			profileID: profileId,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClientProfile(gomock.Any(), gomock.Eq(clientProfile.ID)).
					Times(1).
					Return(db.ClientProfile{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			profileID: "invalidUUID",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClientProfile(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/clients/profiles/%s", tc.profileID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func randomClientProfile() db.ClientProfile {

	return db.ClientProfile{
		ID:      util.RandomUUID(),
		Adm:     "adm:active",
		Kyc:     "kyc:confirmed",
		Name:    util.RandomName(),
		Surname: util.RandomSurname(),
	}
}

func requireBodyMatchClientProfile(t *testing.T, body *bytes.Buffer, clientProfile db.ClientProfile) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotClientProfile db.ClientProfile
	err = json.Unmarshal(data, &gotClientProfile)
	require.NoError(t, err)
	require.Equal(t, clientProfile, gotClientProfile)
}
