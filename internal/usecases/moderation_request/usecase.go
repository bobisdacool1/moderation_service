package moderation_request

import (
	"context"
	"fmt"

	"ModerationService/internal/entity"
)

type (
	Usecase struct {
		service ModerationRequestService
	}

	ModerationRequestService interface {
		CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error)
	}
)

func NewModerationRequestUsecase(service ModerationRequestService) *Usecase {
	return &Usecase{
		service: service,
	}
}

func (u *Usecase) ApproveModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}

func (u *Usecase) DeclineModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}

func (u *Usecase) NextModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	return nil, nil
}

func (u *Usecase) CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	err := u.service.CreateModerationRequest(ctx, request)
	if err != nil {
		return fmt.Errorf("CreateModerationRequest: %w", err)
	}

	return nil
}
