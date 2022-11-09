package api

import (
	"context"

	desc "b2broker/pkg/b2brokerpb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	return &desc.SendMessageResponse{}, nil
}
