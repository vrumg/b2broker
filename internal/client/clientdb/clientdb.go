package clientdb

import (
	"context"
	"fmt"
	"sync"

	"b2broker/internal/model"
)

type Database struct {
	memMap *sync.Map
}

func New() *Database {
	return &Database{
		memMap: &sync.Map{},
	}
}

func (d *Database) RegisterClient(ctx context.Context, clientID string, ch chan model.Message) error {
	_, loaded := d.memMap.LoadOrStore(clientID, ch)
	if !loaded {
		return fmt.Errorf("RegisterClient: client already registered")
	}

	return nil
}

func (d *Database) UnregisterClient(ctx context.Context, clientID string) error {
	d.memMap.Delete(clientID)

	return nil
}

func (d *Database) GetClientData(ctx context.Context, clientID string) (chan model.Message, error) {
	val, ok := d.memMap.Load(clientID)
	if !ok {
		return nil, fmt.Errorf("GetValue: client's data was not found")
	}

	ch, ok := val.(chan model.Message)
	if !ok {
		return nil, fmt.Errorf("GetValue: client's data corrupted")
	}

	return ch, nil
}
