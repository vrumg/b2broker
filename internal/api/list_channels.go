package api

import (
	"context"
	"log"

	desc "b2broker/pkg/b2brokerpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListChannels(ctx context.Context, req *desc.ListChannelsRequest) (*desc.ListChannelsResponse, error) {
	groups, err := i.groupService.GetGroups(ctx)
	if err != nil {
		log.Printf("failed to list channels: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &desc.ListChannelsResponse{List: groups}, nil
}
