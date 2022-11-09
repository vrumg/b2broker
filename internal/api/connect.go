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

	defer func() {
		errDefer := i.clientService.UnregisterClient(stream.Context(), req.Username)
		if errDefer != nil {
			log.Printf("Failed to unregister client: %v", errDefer)
		}
	}()

	for {
		select {
		case _ = <-stream.Context().Done():
			return nil
		case message := <-ch:
			stream.Send(
				&desc.ConnectResponse{
					Data: &desc.MessagePull{
						SenderId: message.SenderID,
						Payload:  message.Payload,
						IsGroup:  message.IsGroup,
					},
				},
			)
		}
	}

	return nil
}
