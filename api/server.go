package api

import (
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for accounting service

type Server struct {
	store  db.Store
	router *gin.Engine
}

// Creates new HTTP Server and setup routings
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	router.POST("/clients/profiles", server.createClientProfile)
	router.GET("/clients/profiles/:id", server.getClientProfile)

	server.router = router
	return server
}

// Start HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
