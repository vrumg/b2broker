package api

import (
	"context"

	desc "b2broker/pkg/b2brokerpb"
)

func (i *Implementation) JoinGroupChat(ctx context.Context, req *desc.JoinGroupChatRequest) (*desc.JoinGroupChatResponse, error) {
	return &desc.JoinGroupChatResponse{}, nil
}
