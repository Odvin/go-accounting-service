package gapi

import (
	"fmt"

	"github.com/Odvin/go-accounting-service/auth"
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/pb"
	"github.com/Odvin/go-accounting-service/util"
)

// Server serves gRPC request for accounting service
type Server struct {
	pb.UnimplementedAccountingServer
	config  util.Config
	auditor auth.Authenticator
	store   db.Store
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	auditor, err := auth.CreatePasetoAuditor(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, auditor: auditor}

	return server, nil
}
