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

	for {
		select {
		case _ = <-stream.Context().Done():
			err := i.clientService.UnregisterClient(stream.Context(), req.Username)
			if err != nil {
				log.Printf("Failed to unregister client: %v", err)
			}
			close(ch)
			return nil

		// No need to check ok - channel is closed in upper case
		case message := <-ch:
			stream.Send(
				&desc.ConnectResponse{
					Data: &desc.Message{
						ReceiverId: message.ReceiverID,
						SenderId:   message.SenderID,
						Payload:    message.Payload,
					},
				},
			)
		}
	}

	return nil
}
