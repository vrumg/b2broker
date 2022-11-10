package api

import (
	"context"
	"log"

	"b2broker/internal/model"
	desc "b2broker/pkg/b2brokerpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	if len(req.Data.SenderId) == 0 {
		log.Printf("failed to send message: received empty sender id")
		return nil, status.Error(codes.InvalidArgument, "sender id is empty")
	}

	if len(req.Data.ReceiverId) == 0 {
		log.Printf("failed to send message: received empty reciever id")
		return nil, status.Error(codes.InvalidArgument, "receiver id is empty")
	}

	// Name of group can overlap with username
	// We check them separately and send message
	isGroupSent := i.sendGroupMessages(ctx, req)
	isUserSent := i.sendUserMessage(ctx, req)

	if !isUserSent && !isGroupSent {
		return nil, status.Error(codes.InvalidArgument, "receiver id is not present")
	}

	return &desc.SendMessageResponse{}, nil
}

func (i *Implementation) sendGroupMessages(ctx context.Context, req *desc.SendMessageRequest) bool {
	var isSent bool

	listeners, err := i.groupService.GetListeners(ctx, req.Data.ReceiverId)
	if err != nil {
		log.Printf("sendGroupMessages: group %s not found, skip to users: %v", req.Data.ReceiverId, err)
		return false
	}

	for val := range listeners {
		err = i.clientService.SendMessage(ctx, model.Message{
			SenderID:   req.Data.ReceiverId, // swap receiver and sender to show channel name
			ReceiverID: val,
			Payload:    req.Data.Payload,
		})
		if err != nil {
			log.Printf("sendGroupMessages: group %s, user %s: %v", req.Data.ReceiverId, val, err)
			continue
		}
		isSent = true
	}

	return isSent
}

func (i *Implementation) sendUserMessage(ctx context.Context, req *desc.SendMessageRequest) bool {
	err := i.clientService.SendMessage(ctx, model.Message{
		SenderID:   req.Data.SenderId,
		ReceiverID: req.Data.ReceiverId,
		Payload:    req.Data.Payload,
	})
	if err != nil {
		log.Printf("sendUserMessage: user %s not found: %v", req.Data.ReceiverId, err)
		return false
	}

	return true
}
