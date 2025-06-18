package moderation_request

import (
	"context"
	"fmt"

	"ModerationService/internal/entity"
)

type (
	ModerationRequestAdapter interface {
		WriteModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error)
	}

	ModerationRequestService struct {
		adapter ModerationRequestAdapter
	}
)

func NewModerationRequestService(adapter ModerationRequestAdapter) *ModerationRequestService {
	return &ModerationRequestService{
		adapter: adapter,
	}
}

func (s *ModerationRequestService) CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	err := s.adapter.WriteModerationRequest(ctx, request)
	if err != nil {
		return fmt.Errorf("create request service: %w", err)
	}

	return nil
}

func (s *ModerationRequestService) GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	return nil, nil
}
