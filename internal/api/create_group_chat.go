package api

import (
	"context"

	desc "b2broker/pkg/b2brokerpb"
)

func (i *Implementation) CreateGroupChat(ctx context.Context, req *desc.CreateGroupChatRequest) (*desc.CreateGroupChatResponse, error) {
	return &desc.CreateGroupChatResponse{}, nil
}
