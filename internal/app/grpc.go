package app

import (
	"context"
	"time"

	"example.com/hasher/gen/proto"
)

// grpc proto implementation
type GrpcHandler struct {
	proto.UnimplementedHasherServer

	hasherService HasherService
}

// Instantiate new handler
func NewGrpcHandler(
	hasherService HasherService,
) *GrpcHandler {
	return &GrpcHandler{
		proto.UnimplementedHasherServer{},
		hasherService,
	}
}

// Get hash
func (g *GrpcHandler) GetHash(
	ctx context.Context,
	request *proto.HashRequest,
) (*proto.HashResponse, error) {
	hash := g.hasherService.GetHash()

	resp := &proto.HashResponse{
		Hash:      hash.Value.String(),
		UpdatedAt: hash.CreatedAt.Format(time.RFC3339),
	}

	return resp, nil
}
