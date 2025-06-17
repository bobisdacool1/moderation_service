package moderation_request

import (
	"context"

	"ModerationService/internal/entity"
)

func (u *UseCase) ApproveModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}
func (u *UseCase) DeclineModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}
func (u *UseCase) NextModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	return nil, nil
}

func (u *UseCase) CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}
