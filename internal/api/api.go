package api

import (
	desc "b2broker/pkg/b2brokerpb"
)

type Implementation struct {
	desc.UnimplementedMessageServiceServer
	clientService clientService
	groupService  groupService
}

// NewAPI return new instance of Implementation.
func NewAPI(
	clientService clientService,
	groupService groupService,
) *Implementation {
	return &Implementation{
		clientService: clientService,
		groupService:  groupService,
	}
}
