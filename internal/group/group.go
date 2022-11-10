package group

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type Service struct {
	memMap     *sync.Map
	listenerMu *sync.Mutex
}

func New() *Service {
	return &Service{
		memMap:     &sync.Map{},
		listenerMu: &sync.Mutex{},
	}
}

func (s *Service) RegisterGroup(ctx context.Context, groupID string) error {
	listeners := make(map[string]struct{}, 0)
	_, loaded := s.memMap.LoadOrStore(groupID, listeners)
	if loaded {
		return fmt.Errorf("RegisterGroup: group already registered")
	}

	return nil
}

func (s *Service) UnregisterGroup(ctx context.Context, groupID string) error {
	s.memMap.Delete(groupID)

	return nil
}

func (s *Service) GetListeners(ctx context.Context, groupID string) (map[string]struct{}, error) {
	val, ok := s.memMap.Load(groupID)
	if !ok {
		return nil, fmt.Errorf("GetListeners: listeners not found")
	}

	listeners, ok := val.(map[string]struct{})
	if !ok {
		return nil, fmt.Errorf("GetListeners: listeners data corrupted")
	}

	return listeners, nil
}

func (s *Service) GetGroups(ctx context.Context) ([]string, error) {
	groupList := make([]string, 0)
	s.memMap.Range(func(key, value any) bool {
		groupName, ok := key.(string)
		if !ok {
			log.Printf("failed to retrieve group name, skip: %v", key)
		}
		groupList = append(groupList, groupName)
		return true
	})

	return groupList, nil
}

func (s *Service) RegisterListener(ctx context.Context, groupID string, clientID string) error {
	val, ok := s.memMap.Load(groupID)
	if !ok {
		return fmt.Errorf("RegisterListener: failed to retrieve group name: %s", groupID)
	}

	listeners, ok := val.(map[string]struct{})
	if !ok {
		return fmt.Errorf("RegisterListener: listeners data corrupted")
	}

	if _, lok := listeners[clientID]; lok {
		return fmt.Errorf("RegisterListener: client is already present in group")
	}

	listeners[clientID] = struct{}{}

	return nil
}

func (s *Service) UnregisterListener(ctx context.Context, groupID string, clientID string) error {
	val, ok := s.memMap.Load(groupID)
	if !ok {
		return fmt.Errorf("UnregisterListener: failed to retrieve group name: %s", groupID)
	}

	listeners, ok := val.(map[string]struct{})
	if !ok {
		return fmt.Errorf("UnregisterListener: listeners data corrupted")
	}

	if _, lok := listeners[clientID]; !lok {
		return fmt.Errorf("UnregisterListener: client is not present in group")
	}

	delete(listeners, clientID)

	return nil
}

func (s *Service) UnregisterListenerFromAllGroups(ctx context.Context, clientID string) error {
	groups, err := s.GetGroups(ctx)
	if err != nil {
		return fmt.Errorf("UnregisterListenerFromAllGroups: failed to get group names")
	}

	for _, groupID := range groups {
		val, ok := s.memMap.Load(groupID)
		if !ok {
			log.Printf("UnregisterListenerFromAllGroups: failed to retrieve group name: %s", groupID)
			continue
		}

		listeners, ok := val.(map[string]struct{})
		if !ok {
			log.Printf("UnregisterListenerFromAllGroups: listeners data corrupted in group %s", groups)
			continue
		}

		if _, lok := listeners[clientID]; !lok {
			log.Printf("UnregisterListenerFromAllGroups: client is not present in group %s", groupID)
			continue
		}

		delete(listeners, clientID)
	}

	return nil
}

func (s *Service) FindGroup(ctx context.Context, groupID string) (map[string]struct{}, error) {
	val, ok := s.memMap.Load(groupID)
	if !ok {
		return nil, fmt.Errorf("GetListeners: listeners not found")
	}

	listeners, ok := val.(map[string]struct{})
	if !ok {
		return nil, fmt.Errorf("GetListeners: listeners data corrupted")
	}

	return listeners, nil
}
