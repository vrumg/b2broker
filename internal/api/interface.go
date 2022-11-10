package api

import (
	"context"

	"b2broker/internal/model"
)

type clientService interface {
	RegisterClient(ctx context.Context, clientID string, ch chan model.Message) error
	UnregisterClient(ctx context.Context, clientID string) error
	SendMessage(ctx context.Context, msg model.Message) error
}

type groupService interface {
	RegisterGroup(ctx context.Context, groupID string) error
	UnregisterGroup(ctx context.Context, groupID string) error
	GetListeners(ctx context.Context, groupID string) (map[string]struct{}, error)
	GetGroups(ctx context.Context) ([]string, error)
	RegisterListener(ctx context.Context, groupID string, clientID string) error
	UnregisterListener(ctx context.Context, groupID string, clientID string) error
	FindGroup(ctx context.Context, groupID string) (map[string]struct{}, error)
	UnregisterListenerFromAllGroups(ctx context.Context, clientID string) error
}
