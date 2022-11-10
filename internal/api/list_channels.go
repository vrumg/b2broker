package api

import (
	"context"
	"log"

	desc "b2broker/pkg/b2brokerpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ListChannels(ctx context.Context, req *desc.ListChannelsRequest) (*desc.ListChannelsResponse, error) {
	groups, err := i.groupService.GetGroups(ctx)
	if err != nil {
		log.Printf("failed to list channels: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := make([]*desc.Group, 0, len(groups))
	for _, group := range groups {
		listeners, err := i.groupService.GetListeners(ctx, group)
		if err != nil {
			log.Printf("failed to get listeners: %v", err)
			continue
		}
		clientIDList := make([]string, 0, len(listeners))
		for val := range listeners {
			clientIDList = append(clientIDList, val)
		}
		resp = append(resp, &desc.Group{
			GroupId:   group,
			Usernames: clientIDList,
		})
	}

	return &desc.ListChannelsResponse{Groups: resp}, nil
}
