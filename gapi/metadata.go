package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	data := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		if userAgents := md.Get("user-agent"); len(userAgents) > 0 {
			data.UserAgent = userAgents[0]
		}

	}

	if p, ok := peer.FromContext(ctx); ok {
		data.ClientIP = p.Addr.String()
	}

	return data
}
