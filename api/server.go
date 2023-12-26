package api

import (
	"fmt"

	"github.com/Odvin/go-accounting-service/auth"
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for accounting service

type Server struct {
	config  util.Config
	auditor auth.Authenticator
	store   db.Store
	router  *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	auditor, err := auth.CreatePasetoAuditor(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, auditor: auditor}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/clients/profiles", server.createClientProfile)
	router.POST("/clients/login", server.createClientToken)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.auditor))
	authRoutes.GET("/clients/profiles", server.getClientProfile)

	server.router = router
}

// Start HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
