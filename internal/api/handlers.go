package api

import (
	"context"

	desc "b2broker/pkg/b2brokerpb"
)

func (i *Implementation) Connect(req *desc.ConnectRequest, stream desc.MessageService_ConnectServer) error {

	return nil
}

func (i *Implementation) JoinGroupChat(ctx context.Context, req *desc.JoinGroupChatRequest) (*desc.JoinGroupChatResponse, error) {
	return &desc.JoinGroupChatResponse{}, nil
}

func (i *Implementation) LeftGroupChat(ctx context.Context, req *desc.LeftGroupChatRequest) (*desc.LeftGroupChatResponse, error) {
	return &desc.LeftGroupChatResponse{}, nil
}

func (i *Implementation) CreateGroupChat(ctx context.Context, req *desc.CreateGroupChatRequest) (*desc.CreateGroupChatResponse, error) {
	return &desc.CreateGroupChatResponse{}, nil
}

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	return &desc.SendMessageResponse{}, nil
}

func (i *Implementation) ListChannels(ctx context.Context, req *desc.ListChannelsRequest) (*desc.ListChannelsResponse, error) {
	return &desc.ListChannelsResponse{}, nil
}
