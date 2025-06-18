package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type (
	HealthcheckHandler struct {
		healthcheckUsecase HealthcheckUsecase
	}

	HealthcheckUsecase interface {
		Healthcheck(ctx context.Context) error
	}
)

func NewHealthcheckHandler(usecase HealthcheckUsecase) *HealthcheckHandler {
	return &HealthcheckHandler{
		healthcheckUsecase: usecase,
	}
}

func (h *HealthcheckHandler) Healthcheck(c *fiber.Ctx) error {
	err := h.healthcheckUsecase.Healthcheck(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Status(fiber.StatusOK)
	return nil
}
