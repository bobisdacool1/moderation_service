package moderation_request

import (
	"context"

	"ModerationService/internal/entity"
)

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
	return nil
}
