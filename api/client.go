package api

import (
	"net/http"

	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateClientProfileRequest struct {
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

func (server *Server) createClientProfile(ctx *gin.Context) {
	var req CreateClientProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateClientProfileParams{
		ID:      util.RandomUUID(),
		Adm:     db.AdministrativeStatusAdmActive,
		Kyc:     db.KycStatusKycConfirmed,
		Name:    req.Name,
		Surname: req.Surname,
	}

	id, err := server.store.CreateClientProfile(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

type GetClientProfileRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getClientProfile(ctx *gin.Context) {
	var req GetClientProfileRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	profileId, _ := uuid.Parse(req.ID)

	profile, err := server.store.GetClientProfile(ctx, profileId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
