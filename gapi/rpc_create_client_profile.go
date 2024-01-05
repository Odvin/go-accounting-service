package gapi

import (
	"context"

	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/pb"
	"github.com/Odvin/go-accounting-service/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateClientProfile(ctx context.Context, req *pb.CreateClientProfileRequest) (*pb.CreateClientProfileResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateClientProfileParams{
		ID:       util.RandomUUID(),
		Adm:      db.AdministrativeStatusAdmActive,
		Kyc:      db.KycStatusKycConfirmed,
		Name:     req.GetName(),
		Surname:  req.GetSurname(),
		Password: hashedPassword,
		Email:    req.GetEmail(),
	}

	id, err := server.store.CreateClientProfile(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "client already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create the client: %s", err)
	}

	res := &pb.CreateClientProfileResponse{
		Id: id.String(),
	}
	return res, nil
}
