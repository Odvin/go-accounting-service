package api

import (
	"errors"
	"fmt"
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
	RefreshToken       string           `json:"refresh_token"`
	RefreshExpired     time.Time        `json:"refresh_token_expired"`
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

	refreshToken, refreshPayload, err := server.auditor.CreateToken(
		profile.ID,
		util.DepositorRole,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:      refreshPayload.ID,
		Sub:     refreshPayload.Sub,
		Refresh: refreshToken,
		Agent:   ctx.Request.UserAgent(),
		Ip:      ctx.ClientIP(),
		Blocked: false,
		Expires: refreshPayload.Expired,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createClientTokenResponse{
		SessionID:          session.ID,
		AccessToken:        accessToken,
		AccessTokenExpired: accessPayload.Expired,
		RefreshToken:       session.Refresh,
		RefreshExpired:     session.Expires,
		ClientInfo:         ClientProfileResponse(profile),
	}

	ctx.JSON(http.StatusOK, rsp)
}

// ================== refreshClientToken ==================

type refreshClientTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type refreshClientTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) refreshClientToken(ctx *gin.Context) {
	var req refreshClientTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshPayload, err := server.auditor.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if session.Blocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.Sub != refreshPayload.Sub {
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.Refresh != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if time.Now().After(session.Expires) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.auditor.CreateToken(
		refreshPayload.Sub,
		refreshPayload.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := refreshClientTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.Expired,
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
