package client

import (
	"context"
	"fmt"

	"b2broker/internal/model"
)

type Service struct {
	repo repo
}

func New(
	repo repo,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterClient(ctx context.Context, clientID string, ch chan model.Message) error {
	err := s.repo.RegisterClient(ctx, clientID, ch)
	if err != nil {
		return fmt.Errorf("RegisterClient: %w", err)
	}

	return nil
}

func (s *Service) UnregisterClient(ctx context.Context, clientID string) error {
	err := s.repo.UnregisterClient(ctx, clientID)
	if err != nil {
		return fmt.Errorf("UnregisterClient: %w", err)
	}

	return nil
}

func (s *Service) GetWriteChan(ctx context.Context, clientID string) (chan<- model.Message, error) {
	ch, err := s.repo.GetClientData(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("GetWriteChan: %w", err)
	}

	return ch, nil
}
