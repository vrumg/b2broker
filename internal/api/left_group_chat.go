package api

import (
	"context"

	desc "b2broker/pkg/b2brokerpb"
)

func (i *Implementation) LeftGroupChat(ctx context.Context, req *desc.LeftGroupChatRequest) (*desc.LeftGroupChatResponse, error) {
	return &desc.LeftGroupChatResponse{}, nil
}
