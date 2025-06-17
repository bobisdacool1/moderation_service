package moderation_request

import (
	"context"

	"ModerationService/internal/entity"
)

type (
	ModerationRequestAdapter interface {
		CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		UpdateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
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
	return nil
}

func (s *ModerationRequestService) ChangeModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}

func (s *ModerationRequestService) GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	return nil, nil
}
