package moderation_request

import (
	"context"

	"ModerationService/internal/entity"
)

type (
	Adapter struct {
	}
)

func NewModerationRequestAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}
func (a *Adapter) UpdateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error {
	return nil
}
func (a *Adapter) GetModerationRequest(ctx context.Context) (*entity.ModerationRequest, error) {
	return nil, nil
}
