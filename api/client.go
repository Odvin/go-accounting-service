package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/Odvin/go-accounting-service/auth"
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type clientPublicInfo struct {
	Name            string    `json:"name"`
	Surname         string    `json:"surname"`
	Email           string    `json:"email"`
	Updated         time.Time `json:"updated"`
	Created         time.Time `json:"created"`
	PasswordUpdated time.Time `json:"password_updated"`
}

func ClientProfileResponse(profile db.ClientProfile) clientPublicInfo {
	return clientPublicInfo{
		Name:            profile.Name,
		Surname:         profile.Surname,
		Email:           profile.Email,
		Updated:         profile.Updated,
		Created:         profile.Created,
		PasswordUpdated: profile.PasswordUpdated,
	}
}

// ================== createClientProfile ==================

type CreateClientProfileRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type CreateClientProfileResponse struct {
	ID string `json:"id"`
}

func (server *Server) createClientProfile(ctx *gin.Context) {
	var req CreateClientProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateClientProfileParams{
		ID:       util.RandomUUID(),
		Adm:      db.AdministrativeStatusAdmActive,
		Kyc:      db.KycStatusKycConfirmed,
		Name:     req.Name,
		Surname:  req.Surname,
		Password: hashedPassword,
		Email:    req.Email,
	}

	id, err := server.store.CreateClientProfile(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := CreateClientProfileResponse{
		ID: id.String(),
	}

	ctx.JSON(http.StatusOK, res)
}

// ================== createClientToken ==================

type createClientTokenRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type createClientTokenResponse struct {
	SessionID          uuid.UUID        `json:"session_id"`
	AccessToken        string           `json:"access_token"`
	AccessTokenExpired time.Time        `json:"access_token_expired"`
	ClientInfo         clientPublicInfo `json:"client"`
}

func (server *Server) createClientToken(ctx *gin.Context) {
	var req createClientTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	profile, err := server.store.GetClientProfileByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, profile.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.auditor.CreateToken(
		profile.ID,
		util.DepositorRole,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createClientTokenResponse{
		SessionID:          uuid.New(),
		AccessToken:        accessToken,
		AccessTokenExpired: accessPayload.Expired,
		ClientInfo:         ClientProfileResponse(profile),
	}

	ctx.JSON(http.StatusOK, rsp)
}

// ================== getClientProfile ==================

func (server *Server) getClientProfile(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*auth.Payload)

	profile, err := server.store.GetClientProfile(ctx, authPayload.Sub)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := ClientProfileResponse(profile)

	ctx.JSON(http.StatusOK, rsp)
}
