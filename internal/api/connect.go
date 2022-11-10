package api

import (
	"log"

	"b2broker/internal/model"
	desc "b2broker/pkg/b2brokerpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Connect(req *desc.ConnectRequest, stream desc.MessageService_ConnectServer) error {
	if len(req.Username) == 0 {
		return status.Error(codes.InvalidArgument, "Username is empty")
	}

	ch := make(chan model.Message)
	err := i.clientService.RegisterClient(stream.Context(), req.Username, ch)
	if err != nil {
		return status.Error(codes.Internal, "Failed to connect")
	}

	log.Printf("client id %s connected", req.Username)

	for {
		select {
		case _ = <-stream.Context().Done():
			err = i.clientService.UnregisterClient(stream.Context(), req.Username)
			if err != nil {
				log.Printf("failed to unregister client: %v", err)
			}
			err = i.groupService.UnregisterListenerFromAllGroups(stream.Context(), req.Username)
			if err != nil {
				log.Printf("failed to unregister client groups: %v", err)
			}
			close(ch)
			return nil

		// No need to check ok - channel is closed in upper case
		case msg := <-ch:
			err = stream.Send(
				&desc.ConnectResponse{
					Data: &desc.Message{
						ReceiverId: msg.ReceiverID,
						SenderId:   msg.SenderID,
						Payload:    msg.Payload,
					},
				},
			)
			if err != nil {
				log.Printf("failed to send message from %s to %s: %v", msg.SenderID, msg.ReceiverID, err)
			}
		}
	}

	log.Printf("Client id %s disconnected", req.Username)

	return nil
}
