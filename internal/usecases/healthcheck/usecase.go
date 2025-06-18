package healthcheckusecase

import (
	"context"
)

type (
	Usecase struct {
		service HealthcheckService
	}

	HealthcheckService interface {
		Ping(ctx context.Context) error
	}
)

func NewHealthcheckUsecase(service HealthcheckService) *Usecase {
	return &Usecase{
		service: service,
	}
}

func (u *Usecase) Healthcheck(ctx context.Context) error {
	return u.service.Ping(ctx)
}
