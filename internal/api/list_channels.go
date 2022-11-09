package api

import (
	"context"

	desc "b2broker/pkg/b2brokerpb"
)

func (i *Implementation) ListChannels(ctx context.Context, req *desc.ListChannelsRequest) (*desc.ListChannelsResponse, error) {
	return &desc.ListChannelsResponse{}, nil
}
