package api

import (
	"context"

	"b2broker/internal/model"
)

type clientService interface {
	RegisterClient(ctx context.Context, clientID string, ch chan model.Message) error
	UnregisterClient(ctx context.Context, clientID string) error
	GetWriteChan(ctx context.Context, clientID string) (chan<- model.Message, error)
}
