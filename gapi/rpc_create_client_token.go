package gapi

import (
	"context"
	"errors"

	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/pb"
	"github.com/Odvin/go-accounting-service/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateClientToken(ctx context.Context, req *pb.CreateClientTokenRequest) (*pb.CreateClientTokenResponse, error) {

	profile, err := server.store.GetClientProfileByEmail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "client not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find the client")
	}

	err = util.CheckPassword(req.GetPassword(), profile.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password")
	}

	accessToken, accessPayload, err := server.auditor.CreateToken(
		profile.ID,
		util.DepositorRole,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create the access token")
	}

	refreshToken, refreshPayload, err := server.auditor.CreateToken(
		profile.ID,
		util.DepositorRole,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create the refresh token")
	}

	mData := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:      refreshPayload.ID,
		Sub:     refreshPayload.Sub,
		Refresh: refreshToken,
		Agent:   mData.UserAgent,
		Ip:      mData.ClientIP,
		Blocked: false,
		Expires: refreshPayload.Expired,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store client session")
	}

	rsp := &pb.CreateClientTokenResponse{
		ClientInfo: &pb.ClientPublicInfo{
			Name:            profile.Name,
			Surname:         profile.Surname,
			Email:           profile.Email,
			PasswordUpdated: timestamppb.New(profile.PasswordUpdated),
			Created:         timestamppb.New(profile.Created),
			Updated:         timestamppb.New(profile.PasswordUpdated),
		},
		SessionId:           session.ID.String(),
		AccessToken:         accessToken,
		RefreshToken:        session.Refresh,
		RefreshTokenExpired: timestamppb.New(session.Expires),
		AccessTokenExpired:  timestamppb.New(accessPayload.Expired),
	}

	return rsp, nil
}
