package moderation_request

import (
	"context"

	"ModerationService/internal/entity"
)

type (
	ModerationRequestService interface {
		CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		ChangeModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error)
	}
)

type (
	ModerationRequestUsecase interface {
		CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		ApproveModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		DeclineModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		NextModerationRequest(ctx context.Context) (*entity.ModerationRequest, error)
	}
)
