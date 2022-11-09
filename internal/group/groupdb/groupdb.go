package groupdb

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type Database struct {
	memMap     *sync.Map
	listenerMu sync.Mutex
}

func New() *Database {
	return &Database{
		memMap: &sync.Map{},
	}
}

// todo rework listeners to map with mutex load/mutate
func (d *Database) RegisterGroup(ctx context.Context, groupID string, listeners map[string]struct{}) error {
	_, loaded := d.memMap.LoadOrStore(groupID, listeners)
	if !loaded {
		return fmt.Errorf("RegisterGroup: group already registered")
	}

	return nil
}

func (d *Database) UnregisterGroup(ctx context.Context, groupID string) error {
	d.memMap.Delete(groupID)

	return nil
}

func (d *Database) GetListeners(ctx context.Context, groupID string) ([]string, error) {
	val, ok := d.memMap.Load(groupID)
	if !ok {
		return nil, fmt.Errorf("GetListeners: listeners not found")
	}

	listeners, ok := val.([]string)
	if !ok {
		return nil, fmt.Errorf("GetListeners: listeners data corrupted")
	}

	return listeners, nil
}

func (d *Database) GetGroups(ctx context.Context) ([]string, error) {
	groupList := make([]string, 0)
	d.memMap.Range(func(key, value any) bool {
		groupName, ok := key.(string)
		if !ok {
			log.Printf("failed to retrieve group name, skip: %v", key)
		}
		groupList = append(groupList, groupName)
		return true
	})

	return groupList, nil
}

func (d *Database) RegisterListener(ctx context.Context, groupID string) error {
	val, ok := d.memMap.Load(groupID)
	if !ok {
		return fmt.Errorf("failed to retrieve group name: %s", groupID)
	}

	return nil
}

func (d *Database) UnregisterListener(ctx context.Context, groupID string) error {

	return nil
}
