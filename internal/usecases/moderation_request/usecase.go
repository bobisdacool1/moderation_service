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
		ReleaseModerationRequest(ctx context.Context, id string) (*entity.ModerationRequest, error)
		CreateDeclinedRequest(ctx context.Context, request *entity.ModerationRequest) error
		CreateApprovedRequest(ctx context.Context, request *entity.ModerationRequest) error
	}
)

func NewModerationRequestUsecase(service ModerationRequestService) *Usecase {
	return &Usecase{
		service: service,
	}
}

func (u *Usecase) ApproveModerationRequest(ctx context.Context, id string) error {
	moderationRequest, err := u.service.ReleaseModerationRequest(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to release moderation request: %w", err)
	}

	moderationRequest.Status = "approved"

	err = u.service.CreateApprovedRequest(ctx, moderationRequest)
	if err != nil {
		return fmt.Errorf("failed to create approved request: %w", err)
	}

	return nil
}

func (u *Usecase) DeclineModerationRequest(ctx context.Context, id string) error {
	moderationRequest, err := u.service.ReleaseModerationRequest(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to release moderation request: %w", err)
	}

	moderationRequest.Status = "declined"

	err = u.service.CreateDeclinedRequest(ctx, moderationRequest)
	if err != nil {
		return fmt.Errorf("failed to create declined request: %w", err)
	}

	return nil
}

func (u *Usecase) NextModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	moderationRequest, err := u.service.GetModerationRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch moderation request: %w", err)
	}

	return moderationRequest, nil
}

func (u *Usecase) CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	err := u.service.CreateModerationRequest(ctx, request)
	if err != nil {
		return fmt.Errorf("CreateModerationRequest: %w", err)
	}

	return nil
}
