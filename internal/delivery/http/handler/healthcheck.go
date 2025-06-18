package handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	HealthcheckHandler struct {
		healthcheckAdapter HealthcheckAdapter
	}

	HealthcheckAdapter interface {
		Ping(ctx context.Context) error
	}
)

func NewHealthcheck() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Healthcheck(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
