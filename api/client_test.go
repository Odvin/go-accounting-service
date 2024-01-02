package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Odvin/go-accounting-service/auth"
	mockdb "github.com/Odvin/go-accounting-service/db/mock"
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type ClientProfileSetup struct {
	Adm      db.AdministrativeStatus
	Kyc      db.KycStatus
	ID       uuid.UUID
	Password string
}

func randomCreateClientProfileRequest() (CreateClientProfileRequest, ClientProfileSetup) {
	password := uuid.New().String()

	hashedPassword, _ := util.HashPassword(password)

	req := CreateClientProfileRequest{
		Name:     util.RandomName(),
		Surname:  util.RandomSurname(),
		Password: password,
		Email:    util.RandomEmail(),
	}

	setup := ClientProfileSetup{
		Adm:      db.AdministrativeStatusAdmActive,
		Kyc:      db.KycStatusKycConfirmed,
		ID:       util.RandomUUID(),
		Password: hashedPassword,
	}

	return req, setup
}

func requireCreateClientProfileResponse(t *testing.T, body *bytes.Buffer, res CreateClientProfileResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var resData CreateClientProfileResponse
	err = json.Unmarshal(data, &resData)
	require.NoError(t, err)
	require.Equal(t, res, resData)
}

func TestCreateClientProfile(t *testing.T) {

	req, setup := randomCreateClientProfileRequest()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":     req.Name,
				"surname":  req.Surname,
				"password": req.Password,
				"email":    req.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateClientProfileParams{
					ID:       setup.ID,
					Adm:      setup.Adm,
					Kyc:      setup.Kyc,
					Name:     req.Name,
					Surname:  req.Surname,
					Password: setup.Password,
					Email:    req.Email,
				}
				store.EXPECT().
					CreateClientProfile(gomock.Any(), gomock.AssignableToTypeOf(arg)).
					Times(1).
					Return(setup.ID, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireCreateClientProfileResponse(t, recorder.Body, CreateClientProfileResponse{ID: setup.ID.String()})
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"name":     req.Name,
				"surname":  req.Surname,
				"password": "123",
				"email":    req.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClientProfile(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"name":     req.Name,
				"surname":  req.Surname,
				"password": req.Password,
				"email":    "invalid-email",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClientProfile(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name":     req.Name,
				"surname":  req.Surname,
				"password": req.Password,
				"email":    req.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClientProfile(gomock.Any(), gomock.Any()).
					Times(1).
					Return(uuid.New(), sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateUsername",
			body: gin.H{
				"name":     req.Name,
				"surname":  req.Surname,
				"password": req.Password,
				"email":    req.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClientProfile(gomock.Any(), gomock.Any()).
					Times(1).
					Return(uuid.New(), db.ErrUniqueViolation)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
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

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/clients/profiles"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestCreateClientToken(t *testing.T) {
	credentials := createClientTokenRequest{
		Email:    "client@email.com",
		Password: "accounting",
	}

	profile := db.ClientProfile{
		ID:       uuid.MustParse("69359037-9599-48e7-b8f2-48393c019135"),
		Adm:      "adm:active",
		Kyc:      "kyc:confirmed",
		Name:     "John",
		Surname:  "Dou",
		Password: "$2a$10$uXnAeTzl0fSialQvRgZ4P.ukoJbJUnsjjY3gtoPYg2q/Mfe4GdB1G",
		Email:    "client@email.com",
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":    credentials.Email,
				"password": credentials.Password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClientProfileByEmail(gomock.Any(), gomock.Eq(credentials.Email)).
					Times(1).
					Return(profile, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/clients/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetClientProfile(t *testing.T) {
	clientProfile := randomClientProfile()
	clientID := uuid.New()

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, auditor auth.Authenticator)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, auditor auth.Authenticator) {
				addAuthorization(t, request, auditor, authorizationTypeBearer, clientID, util.DepositorRole, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClientProfile(gomock.Any(), gomock.Eq(clientID)).
					Times(1).
					Return(clientProfile, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchClientProfile(t, recorder.Body, clientProfile)
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

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/clients/profiles"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.auditor)
			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func randomClientProfile() db.ClientProfile {

	return db.ClientProfile{
		Email:   "client@email.com",
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
