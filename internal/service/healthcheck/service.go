package healthcheckservice

import (
	"context"
)

type (
	HealthcheckAdapter interface {
		Ping(ctx context.Context) error
	}

	HealthcheckService struct {
		adapter HealthcheckAdapter
	}
)

func NewHealthcheckService(adapter HealthcheckAdapter) *HealthcheckService {
	return &HealthcheckService{
		adapter: adapter,
	}
}

func (h *HealthcheckService) Ping(ctx context.Context) error {
	return h.adapter.Ping(ctx)
}
