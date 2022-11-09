package client

import (
	"context"

	"b2broker/internal/model"
)

type repo interface {
	RegisterClient(ctx context.Context, clientID string, ch chan model.Message) error
	UnregisterClient(ctx context.Context, clientID string) error
	GetClientData(ctx context.Context, clientID string) (chan model.Message, error)
}
