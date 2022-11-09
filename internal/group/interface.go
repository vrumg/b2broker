package group

import (
	"context"
)

type repo interface {
	RegisterGroup(ctx context.Context, groupID string, listeners []string) error
	UnregisterGroup(ctx context.Context, groupID string) error
	GetListeners(ctx context.Context, groupID string) ([]string, error)
	GetGroups(ctx context.Context) ([]string, error)
}
