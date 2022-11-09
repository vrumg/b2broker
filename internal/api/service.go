package api

import (
	desc "b2broker/pkg/b2brokerpb"
)

type Implementation struct {
	desc.UnimplementedMessageServiceServer
	clientService clientService
}

// NewAPI return new instance of Implementation.
func NewAPI(
	clientService clientService,
) *Implementation {
	return &Implementation{
		clientService: clientService,
	}
}
