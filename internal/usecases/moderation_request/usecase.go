package moderation_request

import (
	"context"

	"ModerationService/internal/entity"
)

type (
	Usecase struct {
		service ModerationRequestService
	}

	ModerationRequestService interface {
		CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		ChangeModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error)
	}
)

func NewModerationRequestUsecase(service ModerationRequestService) *Usecase {
	return &Usecase{
		service: service,
	}
}
